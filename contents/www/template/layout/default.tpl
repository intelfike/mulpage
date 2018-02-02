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
		<h1>intelfike</h1>
	</a>
</header>

<div id="middle">
	<nav>
		<ul>
			<li><a href="/_top">トップページ</a></li>
			<li><a href="/_interested">興味</a></li>
			<li><a href="/_works">作品</a></li>
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