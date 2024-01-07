

# ENDPOINTS    



* Get: 
    `/generator/download/{name}`                        Download project

* Post:
    `/generator/create/go/cobra`                         Create Cobra Project, pass in a json projectname and modname.
* Post:
    `/generator/create/go/api`                           Create Api Project, pass a json like [here](request.json)



Json values in  `generator/create/go/api`: 

* `project` - string - put anything here
* `mod` - string - put anything here
* `router` - string - It depends on what the user marks. Options: ("chi", "gin", "http", "fiber", "gorilla").
* `db` - string - It depends on what the user marks. Options: ("none", "postgresql", "mysql").
* `ui` - bool - It depends on what the user marks.

