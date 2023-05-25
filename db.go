/*
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-05-26 04:15:12
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-05-26 04:22:31
 * @FilePath: /short/db.go
 * @Description:
 *
 * Copyright (c) 2023 by shgopher, All Rights Reserved.
 */
package short

type DB interface {
	Add(v *Node) (shortURL string, err error) // add v to db
	Delete(shortURL string)                   // delete v
	Change(vi *Node, r string) error          // change v to vi
	Find(shortURL string) (string, error)     // find v from db.
}
