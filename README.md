# 📦 kube-log-collector

**Kube-log-collector is designed to streamline your Kubernetes logging experience, making it as easy as pie!**

Managing logs in a Kubernetes environment can sometimes be confusing. That’s where `kube-log-collector` comes into play! It collects all the logs from a specific namespace or certain pod labels and consolidates them into a single file. You can easily analyze, store logs, and keep your daily logs organized. 🎉

## 🌟 Features
- **One Command Collection**: Gather all your logs in one go.
- **Save as a File**: Keep your logs in a specific file for better organization.
- **Fun and User-Friendly**: Just a couple of options to add, and you’re good to go!

## 🚀 Installation

Getting started with `kube-log-collector` is super easy! Just a few steps:

### Running with Docker

You can run `kube-log-collector` using Docker. Here’s what you need to do:

1. First, build the Docker image:
    ```sh
    docker build -t kube-log-collector .
    ```

2. Then, use the `kube-log-collector` tool to collect logs:
    ```sh
    docker run --rm kube-log-collector -n <namespace> -o <output_file>
    ```

   Specify the namespace with `-n` and the name of the file to save the logs with `-o`. That’s it!

### Running from Source Code

If you prefer a more "hands-on" experience, you can run it from the source code:

1. Make sure you have Go installed: [Go Installation](https://golang.org/doc/install)
2. Clone the repository:
    ```sh
    git clone https://github.com/username/kube-log-collector.git
    cd kube-log-collector
    ```
3. Download the necessary dependencies:
    ```sh
    go mod download
    ```
4. Build the application:
    ```sh
    go build -o kube-log-collector main.go
    ```
5. Run it to collect your logs:
    ```sh
    ./kube-log-collector -n <namespace> -o <output_file>
    ```

## 🎮 Usage

Here are a few examples of how to use `kube-log-collector`:

### Collect All Logs and Save to a File

To collect all pod logs from the `default` namespace and save them to `logs.txt`, simply run:

```sh
docker run --rm kube-log-collector -n default -o logs.txt
```

## 💡 Tips

If you don’t specify an output file with -o, it will create collected-logs.txt by default.
Organize your logs by providing a different file name within the same namespace.

## 📜 License

MIT License. Check the LICENSE file for details.

## 📞 Support

Got questions? Reach out to me via GitHub Issues or contribute! Let’s develop together and conquer logs as a team! 🌍🚀