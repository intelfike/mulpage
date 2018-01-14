<!DOCTYPE html>
<html lang="ja">
<head>
	<meta charset="UTF-8">
	<title>{{.Title}}</title>
	<link rel="stylesheet" href="css/default.css?_={{.Rand}}">
</head>
<body>

<header>
	<h1>isearの解説</h1>
</header>

<nav>
</nav>

<article>
{{template "page" .}}
</article>

<footer>
	&copy;2018.01 intelfike(intelfike@gmail.com)
</footer>

</body>
</html>