# Installation Guide for Einar CLI on MacOS

Follow these steps to install the Einar Command Line Interface (CLI) on MacOS.

## Step 1: Clone the Einar CLI Repository

First, clone the einar-cli repository from GitHub using the following command:

    git clone https://github.com/Ignaciojeria/einar-cli.git

## Step 2: Navigate to the Cloned Repository

Once you've cloned the repository, navigate to the directory using the following command:

    cd einar-cli

## Step 3: Build the Binary

Now, build the binary file using the Go build command. The -o flag allows you to specify the output file name:

    go build -o einar

## Step 4: Add the Binary to Your PATH

To ensure that the system can find your newly created binary, you'll need to add it to your PATH. Since you're using the Zsh shell, you can use the following command:

    echo 'export PATH=$PATH:/path-to-your-einar-directory' >> ~/.zshrc
    source ~/.zshrc

## Step 5: Verify the Installation

Finally, confirm that the installation was successful by checking the version of your Einar CLI:

    einar version

If the Einar CLI has been installed correctly, this command will output the version number of your installed tool.
