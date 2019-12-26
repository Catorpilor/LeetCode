package idgd

import (
	"errors"
	"math/rand"
)

type unordered_map struct {
	set map[int]int
}

func NewUm() unordered_map {
	return unordered_map{set: make(map[int]int)}
}

func (u unordered_map) Exists(val int) bool {
	_, exists := u.set[val]
	return exists
}

func (u unordered_map) Iter() (int, error) {
	for k := range u {
		return k, nil
	}
	return -1, errors.New("empty set")
}

func (u unordered_map) Add(val int) {
	u.set[val]++
}

type RandomizedCollection struct {
	hashset map[int]unordered_map
	res     []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{hashset: make(map[int]unordered_map)}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	if u, exists := this.hashset[val]; !exists {
		u = NewUm()
		u.Add(len(this.res))
		this.res = append(this.res, val)
		return true
	}
	this.hashset[val].Add(len(this.res))
	this.res = append(this.res, val)
	// this.hashset[val] = append(this.hashset[val], len(this.res)-1)
	return false
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if _, exists := this.hashset[val]; !exists {
		return false
	}
	sn, n := len(this.hashset[val]), len(this.res)
	v1 := this.hashset[val]
	pos := v1[sn-1]
	key := this.res[n-1]
	// swap
	this.res[pos], this.res[n-1] = this.res[n-1], this.res[pos]
	v := this.hashset[key]
	v[len(v)-1] = pos
	this.hashset[key] = v
	this.hashset[val] = v1[:sn-1]
	this.res = this.res[:n-1]
	return true

}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.res[rand.Intn(len(this.res))]
}
