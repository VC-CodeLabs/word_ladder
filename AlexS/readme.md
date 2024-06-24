# Code Labber!

Implementation for the "Code Ladder" challenge in C#.

### Setup
Requires .NET 8.0: https://dotnet.microsoft.com/en-us/download/dotnet/8.0
Tested on Visual Studio 2022

Per specifications, inputs can be set in [CodeLabber/Program.cs](CodeLabber/Program.cs) via provided variables in the `Main` function.

### Build

Build the [CodeLabber project](CodeLabber/) via Visual Studio. The "Release" configuration will [perform slightly faster](https://learn.microsoft.com/en-us/visualstudio/debugger/how-to-set-debug-and-release-configurations?view=vs-2022).

Or, from this directory:
```pwsh
dotnet build --configuration Release
```

### Run

After building, simply run:
```pwsh
.\CodeLabber\bin\Release\net8.0\CodeLabber.exe
```
