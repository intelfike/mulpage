<!DOCTYPE html>
<html lang="ja">
<head>
	<meta charset="UTF-8">
	<title>{{.Title}} - isear</title>
	<link rel="stylesheet" href="/css/default.css?_={{.Rand}}">
</head>
<body>

<header>
	<!-- <img src="/image/isear_440×280.png" height="64" alt="isear"> -->
	<h1>
		<a href="/" class="normal">isear document</a>
	</h1>
</header>


<div id="middle">
	<nav>
		<ul>
			<li><a href="/_top/">トップページ</a></li>
			<li><a href="/_usage/">isearの使い方</a></li>
			<li><a href="/_support/">サポート情報</a></li>
		</ul>
	</nav>
	<article>
		<h2>{{.Title}}</h2>
{{template "page" .}}
	</article>
</div>

<footer>
	&copy;2018.01 intelfike(intelfike@gmail.com)
</footer>

</body>
</html>