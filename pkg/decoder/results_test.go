package decoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Results(t *testing.T) {
	filePath := "local.txt"

	t.Run("Test_GroupedResults", func(t *testing.T) {
		r := NewResults(filePath)
		listedWord := "lead"
		r.foundMasculineWord("leader", listedWord)
		r.foundMasculineWord("leadership", listedWord)
		mCount := r.getCount(r.masculineCodedWords)
		assert.Equal(t, 2, mCount)
	})

	t.Run("Test_UngroupedResults", func(t *testing.T) {
		r := NewResults(filePath)
		r.foundFeminineWord("dependable", "depend")
		r.foundFeminineWord("empathetic", "empath")
		fCount := r.getCount(r.feminineCodedWords)
		assert.Equal(t, 2, fCount)
	})

	t.Run("Test_Explain", func(t *testing.T) {
		r := NewResults(filePath)

		listedWord := "lead"
		r.foundMasculineWord("leader", listedWord)
		r.foundMasculineWord("leadership", listedWord)
		r.foundMasculineWord("lead", listedWord)
		r.foundFeminineWord("dependable", "depend")
		r.foundFeminineWord("empathetic", "empath")

		resp := r.Explain()
		assert.Equal(t, filePath, resp.File)
		assert.Equal(t, "masculine", resp.Result)
		assert.Equal(t, []string{"lead"}, getKeys(resp.MasculineWords))
		assert.Contains(t, getKeys(resp.FeminineWords), "depend")
		assert.Contains(t, getKeys(resp.FeminineWords), "empath")
	})
}
