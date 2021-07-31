package sets

import (
	"github.com/softronaut/godx/pkg"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMethods(t *testing.T) {
	t.Run("NewHashSet", func(t *testing.T) {
		s := NewHashSet()
		require.True(t, s.IsEmpty())
		s = NewHashSet(nil, true, 2, "3")
		require.False(t, s.IsEmpty())
		require.Equal(t, 4, s.Size())
	})
	t.Run("Clear", func(t *testing.T) {
		s := NewHashSet(1, 2, 3)
		s.Clear()
		require.True(t, s.IsEmpty())
	})
	t.Run("Add/Remove/Contains", func(t *testing.T) {
		s := NewHashSet()
		require.True(t, s.Add("foo"))
		require.False(t, s.Add("foo"))
		require.True(t, s.Contains("foo"))
		require.True(t, s.Add(42))
		require.True(t, s.Contains(42))
		require.Equal(t, 2, s.Size())
		require.True(t, s.Remove("foo"))
		require.False(t, s.Remove("foo"))
		require.Equal(t, 1, s.Size())
		require.True(t, s.Contains(42))
	})
	t.Run("AddAll/RemoveAll/ContainsAll/RetainAll", func(t *testing.T) {
		s := NewHashSet()
		s.AddAll(42, "foo")
		require.True(t, s.ContainsAll(42, "foo"))
		s.AddAll(42, "foo", "bar")
		require.Equal(t, 3, s.Size())
		s.RemoveAll(1, 2, 3)
		require.Equal(t, 3, s.Size())
		s.RemoveAll(42, "bar", "baz")
		require.Equal(t, 1, s.Size())
		// todo
	})
	t.Run("Any/Every/ForEach/Where", func(t *testing.T) {
		size := 100
		s := NewHashSet(pkg.OrderedIntArray(size)...)
		require.Equal(t, size, s.Size())
		require.True(t, s.Any(func(i interface{}) bool {
			v, ok := i.(int)
			return ok && v%2 == 0
		}))
		require.False(t, s.Any(func(i interface{}) bool {
			return i.(int) > size-1
		}))
		require.True(t, s.Every(func(i interface{}) bool {
			_, ok := i.(int)
			return ok
		}))
		require.False(t, s.Every(func(i interface{}) bool { return i.(int) < size-1 }))
		l := 0
		s.ForEach(func(i interface{}) {
			l += i.(int)
		})
		require.Equal(t, (size-1)*size/2, l)
		require.Equal(t, size/2, s.Where(func(i interface{}) bool {
			v, _ := i.(int)
			return v%2 == 0
		}).Size())
	})
	t.Run("ToArray/Difference/Intersection/Union", func(t *testing.T) {
		s1 := NewHashSet(pkg.OrderedIntArray(3)...)
		require.ElementsMatch(t, []interface{}{0, 1, 2}, s1.ToArray())
		s2 := NewHashSet(pkg.OrderedIntArray(5)[1:]...)
		require.ElementsMatch(t, []interface{}{0}, s1.Difference(s2).ToArray())
		require.ElementsMatch(t, []interface{}{3, 4}, s2.Difference(s1).ToArray())
		require.ElementsMatch(t, []interface{}{1, 2}, s1.Intersection(s2).ToArray())
		require.ElementsMatch(t, []interface{}{0, 1, 2, 3, 4}, s1.Union(s2).ToArray())
	})
	t.Run("Map/Reduce", func(t *testing.T) {
		es := pkg.RandomElements(100)
		set := NewHashSet(es...)
		require.Equal(t, 100, set.Size())
		as := set.Map(func(e interface{}) interface{} {
			return e.(pkg.Element).A
		})
		require.Equal(t, 100, len(as))
		for _, v := range as {
			require.True(t, set.Any(func(e interface{}) bool {
				return v == e.(pkg.Element).A
			}))
		}
	})
	t.Run("String/Join", func(t *testing.T) {
		s := NewHashSet()
		require.Equal(t, "HashSet{}", s.String())
		s = NewHashSet(42)
		require.Equal(t, "HashSet{42}", s.String())
		s = NewHashSet(pkg.OrderedIntArray(3)...)
		require.Equal(t, len("HashSet{0, 1, 2}"), len(s.String()))
		require.Equal(t, len("0 - 1 - 2"), len(s.Join(" - ")))
	})
}

//
// Mock
//
