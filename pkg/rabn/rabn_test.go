package rabn

import (
	"fmt"
	"github.com/femnad/rabn/pkg"
	"github.com/go-yaml/yaml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

const (
	tempFile = "/tmp/rabn/test.yml"
	barKey = "bar"
	barCount = 12
	fooCount = 6
	bazKey = "baz"
	bazCount = 6
	quxKey = "qux"
	quxCount = 1
)

func cleanup(t *testing.T) {
	err := os.Remove(tempFile)
	if err != nil {
		t.Errorf("failage: %s", err)
	}
}

func TestSerialization(t *testing.T) {
	h := pkg.history{}
	h.Items = make(map[string]int)
	h.Items["bar"] = 1
	err := h.serialize(tempFile)
	if err != nil {
		t.Errorf("failage: %s", err)
	}
	defer cleanup(t)
}

func testFail(t *testing.T, err error, msg string, args ...interface{}) {
	if err != nil {
		msg = fmt.Sprintf(msg, args)
		t.Errorf("Test fail: %s %s", err, msg)
	}
}

func TestDeserialization(t *testing.T) {
	inputMap := pkg.history{
		Items: map[string]int{bazKey: bazCount},
		Prefix: "",
	}
	input, err := yaml.Marshal(inputMap)
	testFail(t, err, "Marshalling")

	dir, _ := path.Split(tempFile)
	err = os.MkdirAll(dir, 0700)
	testFail(t, err, "Mkdirall")
	err = ioutil.WriteFile(tempFile, input, 0600)
	testFail(t, err, "Writing file")

	defer cleanup(t)

	h := pkg.history{}
	err = h.deserialize(tempFile)
	testFail(t, err, "Deserializing")

	count, _ := h.Items[bazKey]
	if count != bazCount {
		t.Errorf("failage: count for %s = %d, expected %d", bazKey, count, bazCount)
	}
}

var testHistoryMap = pkg.history{
	Items: map[string]int{barKey: barCount, "foo": fooCount, "baz": bazCount, "qux": quxCount},
}

func initTestHistory(t *testing.T) pkg.history {
	input, err := yaml.Marshal(testHistoryMap)
	testFail(t, err, "Marshalling")
	dir, _ := path.Split(tempFile)
	err = os.MkdirAll(dir, 0700)
	testFail(t, err, "Mkdir all")
	err = ioutil.WriteFile(tempFile, input, 0600)
	testFail(t, err, "Writing file")

	h := pkg.history{historyFile: tempFile}
	err = h.deserialize(tempFile)
	testFail(t, err, "De-serializing")

	return h
}

func TestUpdate(t *testing.T) {
	h := initTestHistory(t)
	err := h.serialize(tempFile)
	testFail(t, err, "Serializing")
	defer cleanup(t)

	pkg.addToHistory(h, bazKey)

	err = h.deserialize(tempFile)
	testFail(t, err, "De-serializing")
	newCount, _ := h.Items[bazKey]
	expectedCount := bazCount + 1
	assert.Equal(t, expectedCount, newCount, "Updated count isn't correct")

}

func strippingAssert(t *testing.T, path, expected string, components int) {
	actual := pkg.stripOutput(path, components)
	if expected != actual {
		t.Errorf("test fail in stripping, expected %s, actual %s", expected, actual)
	}
}

func TestStripping(t *testing.T) {
	strippingAssert(t, "/foo/bar/baz", "/foo/bar/baz", 0)

	strippingAssert(t, "/foo/bar/baz", "baz", 1)

	strippingAssert(t, "/foo/bar/baz", "bar/baz", 2)

	strippingAssert(t, "/foo/bar/baz", "/foo/bar/baz", 3)
}

func TestSorting(t *testing.T) {
	h := initTestHistory(t)

	items := pkg.getOrderedItems(h)
	assert.Equal(t, len(testHistoryMap.Items), len(items), "sorted items should be equal length to history map")

	assert.Equal(t, barKey, items[0], "Incorrect first item")
	assert.Equal(t, quxKey, items[len(items) - 1], "Incorrect last item")
}
