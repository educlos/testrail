## Go implementation of the [TestRail](http://www.gurock.com/testrail/) API


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
