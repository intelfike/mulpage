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
{{ .Package.Name }}
{{ if ne .Info.Method "Index" }}
	&gt;{{ .Title }}
{{ end }}
</h2>


<div id="middle">
	<nav>
		<ul>
{{ range $pack := .Packages }}
	{{ if eq $.Info.Package .Key }}
			<li>{{ .Name }}</li>
	{{ else }}
			<li><a href="/_{{ .Key }}/">{{ .Name }}</a></li>
	{{ end }}

	{{ if ne (len .Articles) 1 }}
			<ul>
		{{ range $atc := .ArticlesArray }}
			{{ if ne .Name "" }}
				{{ if and (eq $.Info.Package $pack.Key) (eq $.Template .Key) }}
					<li>{{ .Name }}</li>
				{{ else }}
					<li><a href="/_{{ $pack.Key }}/{{ .Key }}/">{{ .Name }}</a></li>
				{{ end }}
			{{ end }}
		{{ end }}
			</ul>
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