- public methods start with upper-case letter (they 'get exported'), private methods with lower-case letters

- if test is in the same module as production code, we don' t need import and package prefix to access methods

- in imports we either use 'global' packages (simply by name), or relative paths to local packages

- as for interfaces, don' t consume the default examples, instead see here: http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

- t.Error("Expected", expected, "got", actual) => we can pass as many arguments as we want, they will be chained together to a nice error message