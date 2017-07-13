<html>

	<head>
	<title>首页</title>
	<link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
	</head>
	<body>
	
<nav class="navbar navbar-default" role="navigation">
    <div class="container-fluid">
    <div class="navbar-header">
        <a class="navbar-brand" href="#">我的博客</a>
    </div>
    <div>
        <ul class="nav navbar-nav">
            <li class="active"><a href="#">文章</a></li>
            <li><a href="#">照片</a></li>
            
        </ul>
    </div>
    </div>
</nav>

<div class="container">
<ul class="list-group">
{{range $k,$v := .article}}
    <li class="list-group-item">
	标题:<h2>{{.Title}}</h2>
	内容:<p>{{.Content}}</p>
	作者:<p>{{.Author}}</p>
时间:<p>{{date_mh .CreateTime}}</p>
	</li>
 
   {{end}}
</ul>
</div>

	</body>
	<script src="https://cdn.bootcss.com/jquery/2.1.1/jquery.min.js"></script>
	<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</html>