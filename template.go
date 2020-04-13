// Code generated by generate.go DO NOT EDIT.
package main

const (
        tmplStr = `<!DOCTYPE html>
<html lang="en">
  <meta
    name="go-import"
    content="go.seankhliao.com/{{ .Repo }}
        git https://github.com/seankhliao/{{ .Repo }}"
  />
  <meta
    name="go-source"
    content="go.seankhliao.com/{{ .Repo }}
        https://github.com/seankhliao/{{ .Repo }}
        https://github.com/seankhliao/{{ .Repo }}/tree/master{/dir}
        https://github.com/seankhliao/{{ .Repo }}/blob/master{/dir}/{file}#L{line}"
  />
  <meta http-equiv="refresh" content="5;url=https://godoc.org/go.seankhliao.com/{{ .Repo }}" />
  <title>go.seankhliao.com/{{ .Repo }}</title>
  <p>
    source:
    <a
      href="https://github.com/seankhliao/{{ .Repo }}"
      ping="https://statslogger.seankhliao.com/api?trigger=ping&src=go.seankhliao.com/{{ .Repo }}&dst=github.com/seankhliao/{{ .Repo }}"
    >
      github</a
    >
  </p>
  <p>
    docs:
    <a
      href="https://godoc.org/go.seankhliao.com/{{ .Repo }}"
      ping="https://statslogger.seankhliao.com/api?trigger=ping&src=go.seankhliao.com/{{ .Repo }}&dst=godoc.org/go.seankhliao.com/{{ .Repo }}"
    >
      godoc</a
    >
  </p>
  <script>
    let ts0 = new Date();
    let dst = "";
    document.querySelectorAll("a").forEach((el) => {
      el.addEventListener("click", (e) => {
        dst = e.target.href.replace(/(^\w+:|^)\/\//, "");
      });
    });
    window.addEventListener("unload", () => {
      ts1 = new Date();
      navigator.sendBeacon(
        `+"`"+`https://statslogger.seankhliao.com/api?trigger=beacon&src=go.seankhliao.com/{{ .Repo }}&dst=${dst}&dur=${
          ts1 - ts0
        }ms`+"`"+`
      );
    });
  </script>
</html>
`
)