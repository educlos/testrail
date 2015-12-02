## Go implementation of the [TestRail](http://www.gurock.com/testrail/) API

[![GoDoc](https://godoc.org/github.com/Etienne42/testrail?status.svg)](https://godoc.org/github.com/Etienne42/testrail)

Example usage
-------------

```
  client := NewClient("https://example.testrail.com", "username@example.com", "password")

  projectID := 1
  suiteID := 11
  cases, err := client.GetCases(projectID, suiteID)
```


License
-------

MIT
