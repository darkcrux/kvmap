// Package kvmap is a utility to convert an interface{} into a map[string]interface{} relative to the given key.
//
// For non-map, non-slice, non-array, and non-struct data, this simply returns a new map with they given key and
// value as it's only entry.
//
// eg.
//
//     m := kvmap.ToKV("key", "value")
//
//     // this returns map[string]interface{}{ "key": "value" }
//
// For structs, the fields in the data gets traversed and are converted into a map. An optional tag can also be
// specified to define what the keys will be.
//
// eg.
//
//     type Person struct {
//	       Name string `kv:"name"`
//         Age  int    `kv:"age"`
//     }
//
//     func main() {
//	       mrfoo := &Person{
//	           Name: "Mr. Foo",
//             Age:  43,
//         }
//         m := kvmap.ToKV("persons/foo", mrfoo)
//
//         // This returns the following map:
//         // map[string]interface{}{
//         //     "persons/foo/name": "Mr. Foo",
//         //     "persons/foo/age":  43,
//         // }
//     }
//
// For maps, the maps gets traversed and converted into a map[string]interface{} with the data's keys appened to
// the key parameter provided.
//
// eg.
//
//     myMap := map[string]string{
//       "one": "1",
//       "two": "2",
//     }
//     m := kvmap.ToKV("demo", myMap)
//
//     // This returns the following map:
//     // map[string]interface{}{
//	   //     "demo/one": "1",
//     //     "demo/two": "2",
//     // }
//
// For arrays and slices, they get traversed and the element's index are used as the key.
//
// eg.
//
//     mySlice := []string{
//	       "ready",
//         "set",
//         "go",
//     }
//     m := kvmap.ToKV("count", mySlice)
//
//     // This returns the following map:
//     // map[string]interface{}{
//     //     "count/0": "ready",
//     //     "count/1": "set",
//     //     "count/2": "go",
//     // }
//
// For channels, on an empty key is returned.
//
// Nested Data will be processed recursively.
//
// eg.
//
//     type Person struct {
//	       Name     string   `kv:"name"`
//         Age      int      `kv:"age"`
//         Children []string `kv:"children"`
//     }
//
//     func main() {
//	       mrfoo := &Person{
//	           Name:     "Mr. Foo",
//             Age:      43,
//             Children: []string{"fuu", "bar"},
//         }
//         m := kvmap.ToKV("persons/foo", mrfoo)
//
//         // This returns the following map:
//         // map[string]interface{}{
//         //     "persons/foo/name":       "Mr. Foo",
//         //     "persons/foo/age":        43,
//         //     "persons/foo/children/0": "fuu",
//         //     "persons/foo/children/1": "bar",
//         // }
//     }
//
package kvmap
