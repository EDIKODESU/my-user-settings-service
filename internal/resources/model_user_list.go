/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UserList struct {
	Key
	Attributes *UserListAttributes `json:"attributes,omitempty"`
}
type UserListResponse struct {
	Data     UserList `json:"data"`
	Included Included `json:"included"`
}

type UserListListResponse struct {
	Data     []UserList `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustUserList - returns UserList from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUserList(key Key) *UserList {
	var userList UserList
	if c.tryFindEntry(key, &userList) {
		return &userList
	}
	return nil
}
