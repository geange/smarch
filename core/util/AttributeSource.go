package util

import "reflect"

// An AttributeSource contains a list of different AttributeImpls, and methods to add and
// get them. There can only be a single instance of an attribute in the same AttributeSource
// instance. This is ensured by passing in the actual type of the Attribute (Class&lt;Attribute&gt;)
// to the #addAttribute(Class), which then checks if an instance of that type is already
// present. If yes, it returns the instance, otherwise it creates a new instance and returns it.
type AttributeSource struct {

	// These two maps must always be in sync!!!
	// So they are private, final and read-only from the outside (read-only iterators)
	attributes     map[reflect.Type]AttributeImpl
	attributeImpls map[reflect.Type]AttributeImpl
	currentState   []State
}

// State This class holds the state of an AttributeSource.
type State struct {
	attribute AttributeImpl
	next      *State
}

func (s *State) Clone() *State {
	var next *State
	if s.next != nil {
		next = s.next.Clone()
	}
	return &State{
		attribute: s.attribute.Clone(),
		next:      next,
	}
}
