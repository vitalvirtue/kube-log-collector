package collector

import (
    "context"
    "fmt"
    "io"
    "os"

    "github.com/vitalvirtue/kube-log-collector/internal/kubernetes"
    "github.com/vitalvirtue/kube-log-collector/pkg/types"
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Collector struct'ı, log toplama işlemlerini yönetir
type Collector struct {
    client  kubernetes.ClientInterface
    options types.CollectorOptions
}

// NewCollector, yeni bir Collector örneği oluşturur
func NewCollector(client kubernetes.ClientInterface, options types.CollectorOptions) *Collector {
    return &Collector{
        client:  client,
        options: options,
    }
}

// Collect, pod loglarını toplar ve belirtilen dosyaya yazar
func (c *Collector) Collect() error {
    fmt.Println("Starting log collection process")

    pods, err := c.getPods()
    if err != nil {
        fmt.Printf("Error getting pods: %v\n", err)
        return fmt.Errorf("error getting pods: %w", err)
    }
    fmt.Printf("Found %d pods\n", len(pods))

    file, err := os.Create(c.options.OutputFile)
    if err != nil {
        fmt.Printf("Error creating output file: %v\n", err)
        return fmt.Errorf("error creating output file: %w", err)
    }
    defer file.Close()

    for _, pod := range pods {
        if err := c.collectPodLogs(pod, file); err != nil {
            fmt.Printf("Error collecting logs for pod %s: %v\n", pod.Name, err)
            return fmt.Errorf("error collecting logs for pod %s: %w", pod.Name, err)
        }
    }

    fmt.Println("Log collection process completed successfully")
    return nil
}

// collectPodLogs, belirli bir pod için logları toplar ve yazar
func (c *Collector) collectPodLogs(pod corev1.Pod, file *os.File) error {
    fmt.Printf("Collecting logs for pod: %s/%s\n", pod.Namespace, pod.Name)
    ctx := context.Background()
    podLogOptions := &corev1.PodLogOptions{}
    req := c.client.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, podLogOptions)
    podLogs, err := req.Stream(ctx)
    if err != nil {
        fmt.Printf("Error getting log stream for pod %s/%s: %v\n", pod.Namespace, pod.Name, err)
        return err
    }
    defer podLogs.Close()

    _, err = file.WriteString(fmt.Sprintf("--- Logs for pod %s/%s ---\n", pod.Namespace, pod.Name))
    if err != nil {
        fmt.Printf("Error writing log header for pod %s/%s: %v\n", pod.Namespace, pod.Name, err)
        return err
    }

    buffer := make([]byte, 4096)
    for {
        n, err := podLogs.Read(buffer)
        if err != nil && err != io.EOF {
            fmt.Printf("Error reading logs for pod %s/%s: %v\n", pod.Namespace, pod.Name, err)
            return err
        }
        if n == 0 {
            break
        }

        _, err = file.Write(buffer[:n])
        if err != nil {
            fmt.Printf("Error writing logs for pod %s/%s: %v\n", pod.Namespace, pod.Name, err)
            return err
        }
    }

    _, err = file.WriteString("\n\n")
    if err != nil {
        fmt.Printf("Error writing log footer for pod %s/%s: %v\n", pod.Namespace, pod.Name, err)
        return err
    }

    fmt.Printf("Successfully collected logs for pod: %s/%s\n", pod.Namespace, pod.Name)
    return nil
}

// getPods, belirtilen namespace veya etiket için podları getirir
func (c *Collector) getPods() ([]corev1.Pod, error) {
    fmt.Printf("Getting pods for namespace: %s, label: %s\n", c.options.Namespace, c.options.PodLabel)
    ctx := context.Background()
    var listOptions metav1.ListOptions
    if c.options.PodLabel != "" {
        listOptions.LabelSelector = c.options.PodLabel
    }
    podList, err := c.client.CoreV1().Pods(c.options.Namespace).List(ctx, listOptions)
    if err != nil {
        fmt.Printf("Error listing pods: %v\n", err)
        return nil, err
    }
    fmt.Printf("Found %d pods\n", len(podList.Items))
    return podList.Items, nil
}