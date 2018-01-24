<!DOCTYPE html>
<html lang="ja">
<head>
	<meta charset="UTF-8">
	<title>{{ .Title }} | isear</title>
	<link rel="stylesheet" href="/css/default.css?_={{ .Rand }}">
</head>
<body>

<header>
	<a href="/" class="normal" style="padding:10px;">
		<img src="/data/image/isear_document_title.png" height="40" alt="isear">
	</a>
</header>

<h2>
{{ if eq .Info.Method "Index" }}
	{{ .Package.Name }}
{{ else }}
	<a href="/_{{ .Info.Package }}">{{ .Package.Name }}</a>&gt;{{ .Title }}
{{ end }}
</h2>


<div id="middle">
	<nav>
		<ul>
{{ range .Packages }}
	{{ if eq $.Info.Package .Package }}
			<li>{{ .Text }}</li>
	{{ else }}
			<li><a href="/_{{ .Package }}/">{{ .Text }}</a></li>
	{{ end }}
{{ end }}
		</ul>
	</nav>
	<article>
{{ template "page" . }}
	</article>
</div>

<footer>
	&copy;2018.01 intelfike(intelfike@gmail.com)
</footer>

</body>
</html>