# ip-to-int

* ## Build

  * ```bash
    go env -w GO111MODULE=auto
    ```

  * ```bash
    sh build.sh
    ```

* ## Usage

  * ### cd to ./output/bin

    * ```bash
      cd ./output/bin
      ```

  * ### Show help

    * ```bash
      ./ip_to_int -h
      ```

  * ### Convert IP format to int

    * #### Support IPv4

      * ```bash
        ./ip_to_int -a 192.168.0.1
        ```

    * #### Support IPv6

      * ```bash
        ./ip_to_int -a ABCD:0ABC:00AB:000A:0000:0000:0000:1000
        ```

      * ```bash
        ./ip_to_int -a ABCD:ABC:AB:A::1000
        ```