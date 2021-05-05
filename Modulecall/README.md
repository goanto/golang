Edit the example.com/hello module to use your local example.com/greetings module.

For production use, you’d publish the example.com/greetings module from its repository (with a module path that reflected its published location), where Go tools could find it to download it. For now, because you haven't published the module yet, you need to adapt the example.com/hello module so it can find the example.com/greetings code on your local file system. 

$ go mod edit -replace=example.com/greetings=../greetings

From the command prompt in the hello directory, run the go mod tidy command to synchronize the example.com/hello module's dependencies, adding those required by the code, but not yet tracked in the module. 

go mod tidy

after tify in go.mod
require example.com/module v0.0.0-00010101000000-000000000000

 The command found the local code in the greetings directory, then added a require directive to specify that example.com/hello requires example.com/greetings. You created this dependency when you imported the greetings package in hello.go.

The number following the module path is a pseudo-version number -- a generated number used in place of a semantic version number (which the module doesn't have yet).

To reference a published module, a go.mod file would typically omit the replace directive and use a require directive with a tagged version number at the end