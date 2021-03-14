## Etisbew

A simple and opinionated static site generator.

Input directory structure:

```
yoursite/
├╴templates/
│ └╴default.html
├╴blog/
│ └╴my-article.md
├╴horses/
│ └╴american-standardbred.md
├╴index.md
└╴about.md
```

Output structure:

```
yoursite/
└╴output/
  ├╴blog/
  │ └╴my-article/index.html
  ├╴horses/
  │ └╴american-standardbred/index.html
  ├╴index.html
  └╴about/index.html
```

Input files can be either plain markdown or contain an INI header
delimited with `~~~`, but they must use the `.md` extension:

```
index.md
~~~
Title = Welcome
Template = custom.html
~~~

Regular markdown goes here
```

Valid keys in the header are `Title`, `Body`, and `Template`
(default value `default.html`). If you specify `Body` in the header
it will override the contents of the markdown. Not sure why that
would be useful, but it's how I wrote the input parser. :)

Templates use go's [`html/template` package][templates]. Only a
`default.html` template is required, but all templates must be in the
`templates` directory and use the `.html` extention.

```
templates/default.html
<!doctype html>
<meta charset="utf-8">
<title>{{.Title}}</title>

<body>
  {{if .Title}}
  <h2>{{.Title}}</h2>
  {{end}}

  {{.Body}}
</body>
```

And that's pretty much it.

[templates]: https://golang.org/pkg/html/template/
