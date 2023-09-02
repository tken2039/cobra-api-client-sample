# cobra-api-client-sample

## Sample

### Run

```shell
# Plane execution (Table format)
❯ go run ./main.go apicall
User ID                  Email              Name                 Given Name Family Name Nickname
583c3ac3f38e84297c002546 test@test.com      test@test.com        Hello      Test       test
583c5484cb79a5fe593425a9 test1@test.com     test1@test.com       Hello1     Test1      test1
583c57672c7686377d2f66c9 aaa@aaa.com        aaa@aaa.com          John       Dough      aaa
5840b954da0529cd293d76fe a@a.com            a@a.com              Jane       Dough      a
584a9d13e808bcf75f05f580 test9999@test.com                       Dummy      User

# Table format with full data
❯ go run ./main.go apicall --format table --full
User ID                  Email              Name                 Given Name Family Name Nickname   Last IP         Logins Count    Created At                     Updated At                     Last Login                     Email Verified
583c3ac3f38e84297c002546 test@test.com      test@test.com        Hello      Test       test       94.121.163.63   15              2016-11-28T14:10:11.338Z       2016-12-02T01:17:29.310Z       2016-12-02T01:17:29.310Z       true
583c5484cb79a5fe593425a9 test1@test.com     test1@test.com       Hello1     Test1      test1      94.121.168.53   1               2016-11-28T16:00:04.209Z       2016-11-28T16:00:47.203Z       2016-11-28T16:00:47.203Z       true
583c57672c7686377d2f66c9 aaa@aaa.com        aaa@aaa.com          John       Dough      aaa        94.121.168.53   2               2016-11-28T16:12:23.777Z       2016-11-28T16:12:52.353Z       2016-11-28T16:12:52.353Z       true
5840b954da0529cd293d76fe a@a.com            a@a.com              Jane       Dough      a          94.121.163.63   3               2016-12-01T23:59:16.473Z       2016-12-01T23:59:53.474Z       2016-12-01T23:59:53.474Z       true
584a9d13e808bcf75f05f580 test9999@test.com                       Dummy      User                                  0               2016-12-09T12:01:23.787Z       2016-12-09T12:01:23.787Z                                      false

# JSON format
❯ go run ./main.go apicall --format json
[{"user_id":"583c3ac3f38e84297c002546","email":"test@test.com","name":"test@test.com","given_name":"Hello","family_name":"Test","nickname":"test"},{"user_id":"583c5484cb79a5fe593425a9","email":"test1@test.com","name":"test1@test.com","given_name":"Hello1","family_name":"Test1","nickname":"test1"},{"user_id":"583c57672c7686377d2f66c9","email":"aaa@aaa.com","name":"aaa@aaa.com","given_name":"John","family_name":"Dough","nickname":"aaa"},{"user_id":"5840b954da0529cd293d76fe","email":"a@a.com","name":"a@a.com","given_name":"Jane","family_name":"Dough","nickname":"a"},{"user_id":"584a9d13e808bcf75f05f580","email":"test9999@test.com","name":"","given_name":"Dummy","family_name":"User","nickname":""}]

# JSON format with full data
❯ go run ./main.go apicall --format json --full
[{"user_id":"583c3ac3f38e84297c002546","email":"test@test.com","name":"test@test.com","given_name":"Hello","family_name":"Test","nickname":"test","last_ip":"94.121.163.63","logins_count":15,"created_at":"2016-11-28T14:10:11.338Z","updated_at":"2016-12-02T01:17:29.310Z","last_login":"2016-12-02T01:17:29.310Z","email_verified":true},{"user_id":"583c5484cb79a5fe593425a9","email":"test1@test.com","name":"test1@test.com","given_name":"Hello1","family_name":"Test1","nickname":"test1","last_ip":"94.121.168.53","logins_count":1,"created_at":"2016-11-28T16:00:04.209Z","updated_at":"2016-11-28T16:00:47.203Z","last_login":"2016-11-28T16:00:47.203Z","email_verified":true},{"user_id":"583c57672c7686377d2f66c9","email":"aaa@aaa.com","name":"aaa@aaa.com","given_name":"John","family_name":"Dough","nickname":"aaa","last_ip":"94.121.168.53","logins_count":2,"created_at":"2016-11-28T16:12:23.777Z","updated_at":"2016-11-28T16:12:52.353Z","last_login":"2016-11-28T16:12:52.353Z","email_verified":true},{"user_id":"5840b954da0529cd293d76fe","email":"a@a.com","name":"a@a.com","given_name":"Jane","family_name":"Dough","nickname":"a","last_ip":"94.121.163.63","logins_count":3,"created_at":"2016-12-01T23:59:16.473Z","updated_at":"2016-12-01T23:59:53.474Z","last_login":"2016-12-01T23:59:53.474Z","email_verified":true},{"user_id":"584a9d13e808bcf75f05f580","email":"test9999@test.com","name":"","given_name":"Dummy","family_name":"User","nickname":"","last_ip":"","logins_count":0,"created_at":"2016-12-09T12:01:23.787Z","updated_at":"2016-12-09T12:01:23.787Z","last_login":"","email_verified":false}]

❯
```

### Build and run

```shell
❯ make build
Building...

❯ ./bin/cobrapi apicall
User ID                  Email              Name                 Given Name Family Name Nickname
583c3ac3f38e84297c002546 test@test.com      test@test.com        Hello      Test       test
583c5484cb79a5fe593425a9 test1@test.com     test1@test.com       Hello1     Test1      test1
583c57672c7686377d2f66c9 aaa@aaa.com        aaa@aaa.com          John       Dough      aaa
5840b954da0529cd293d76fe a@a.com            a@a.com              Jane       Dough      a
584a9d13e808bcf75f05f580 test9999@test.com                       Dummy      User
```
