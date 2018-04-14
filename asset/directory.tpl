<!DOCTYPE html>
<html>
<head>
  <title>{{ .Root }}</title>
</head>
<body>
  <div class="wrapper">
    <h1>{{ .Root }}</h1>
    <ul>
      {{ range .Directories }}
      <li>{{ .Name }}</li>
      {{ end }}
    </ul>
  </div>
</body>
</html>
