// START OMIT
package handlers

import "net/http"

func Simple(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(loremIpsum))
}
// END OMIT
const loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus hendrerit luctus leo. Mauris feugiat pretium sem eget semper. Donec venenatis, erat ut aliquam sagittis, tortor est condimentum arcu, sit amet dignissim lorem massa in quam. Maecenas sollicitudin lacus augue, eget eleifend eros ultricies sit amet. Curabitur quis tortor consectetur, tempus velit non, condimentum nisi. Donec ultrices viverra dolor. Maecenas nec leo sodales, congue dui eget, dignissim nibh. Interdum et malesuada fames ac ante ipsum primis in faucibus. Fusce eu metus id magna maximus fermentum ut in tellus."
