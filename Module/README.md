
    Create a module -- Write a small module with functions you can call from another module.
    Call your code from another module -- Import and use your new module.
    Return and handle an error -- Add simple error handling.
    Return a random greeting -- Handle data in slices (Go's dynamically-sized arrays).
    Return greetings for multiple people -- Store key/value pairs in a map.
    Add a test -- Use Go's built-in unit testing features to test your code.
    Compile and install the application -- Compile and install your code locally.


go mod init example.com/module

Run the go mod init command, giving it your module path -- here, use example.com/greetings. If you publish a module, this must be a path from which your module can be downloaded by Go tools. That would be your code's repository. 

The go mod init command creates a go.mod file to track your code's dependencies. So far, the file includes only the name of your module and the Go version your code supports. But as you add dependencies, the go.mod file will list the versions your code depends on. This keeps builds reproducible and gives you direct control over which module versions to use. 