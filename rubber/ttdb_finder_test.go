package rubber

import (
	"testing"
	"github.com/PuerkitoBio/goquery"
	"os"
	"gopkg.in/check.v1"
	"sort"
)

func Test(t *testing.T) { check.TestingT(t) }

type MySuite struct{}

var _ = check.Suite(&MySuite{})

func testNewDocument(url string) (*goquery.Document, error) {
	switch url {
	case "http://www.tabletennisdb.com/rubber":
		file, _ := os.Open("test_data/root.html")
		return goquery.NewDocumentFromReader(file)
	case "http://www.tabletennisdb.com/rubber/rubber1.html":
		file, _ := os.Open("test_data/rubber1.html")
		return goquery.NewDocumentFromReader(file)
	case "http://www.tabletennisdb.com/rubber/rubber2.html":
		file, _ := os.Open("test_data/rubber2.html")
		return goquery.NewDocumentFromReader(file)
	}
	return nil, nil
}

type ByName []*Rubber

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name > a[j].Name }

func (s *MySuite) TestFetchRubbersFromDocument(c *check.C) {
	expectedRubbers := []*Rubber{
		{Name: "Donic Acuda P2", 	Speed: 9.1, Spin: 9.7, Control: 9.2, Tackiness: 3.2, Weight: 3.9, SpongeHardness: 4.0, Gears: 8.4, ThrowAngle: 4.4, Consistency: 9.5, Durability: 7.4},
		{Name: "Butterfly Tenergy 05", 	Speed: 9.3, Spin: 9.4, Control: 8.4, Tackiness: 2.3, Weight: 6.8, SpongeHardness: 6.2, Gears: 8.9, ThrowAngle: 7.4, Consistency: 9.4, Durability: 8.1},
	}
	finder := NewTTDBRubberFinder()
	finder.newDocument = testNewDocument

	actualRubbers, err := finder.FindRubbers()

	if err != nil {
		c.Error(err)
	}
	if len(actualRubbers) != 2 {
		c.Fatalf("Should have fetched 2 rubbers, but fetched %d", len(actualRubbers))
	}
	var rubbersByName ByName = actualRubbers
	sort.Sort(rubbersByName)
	c.Check(actualRubbers, check.DeepEquals, expectedRubbers)
}
