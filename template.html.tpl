<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width">
		<title>Moon Demo</title>
		<link rel="shortcut icon" href="//{{.Static}}/favicon.ico">
		{{if not .Hot}}<link rel="stylesheet" type="text/css" href="//{{.Style}}" media="all" />{{end}}
	</head>
	<body>
		<div id="app"></div>
		<script type="text/javascript" src="//{{.Js}}"></script>
	</body>
</html>
