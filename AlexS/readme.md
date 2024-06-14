# Code Labber!

Implementation for the "Code Ladder" challenge in C#.

### Setup
Requires .NET 8.0: https://dotnet.microsoft.com/en-us/download/dotnet/8.0
Tested on Visual Studio 2022

Per specifications, inputs can be set in [CodeLabber/Program.cs](CodeLabber/Program.cs) via provided variables in the `Main` function.

### Build

Build the [CodeLabber project](CodeLabber/) via Visual Studio.

Or, from this directory:
```pwsh
dotnet build
```

### Run

After building, simply run:
```pwsh
.\CodeLabber\bin\Debug\net8.0\CodeLabber.exe
```

Alternatively:
```pwsh
# This will also build the project (slower)
dotnet run --project .\CodeLabber\
```
