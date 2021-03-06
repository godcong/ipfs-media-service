package ffmpeg

import "testing"

// TestRun ...
func TestRun(t *testing.T) {
	b, err := Run("-version")
	t.Log(string(b), err)
}

// TestSpliteMp4 ...
func TestSplitMp4(t *testing.T) {
	b, err := SplitToTS("./tmp/ELTbmjn2IZY6EtLFCibQPL4pIyfMvN8jQS67ntPlFaFo3NUkM3PpCFBgMivKk67W_out.mp4", "./split/ELTbmjn2IZY6EtLFCibQPL4pIyfMvN8jQS67ntPlFaFo3NUkM3PpCFBgMivKk67W")
	t.Log(string(b), err)
}

// TestSplitWithKey ...
func TestSplitWithKey(t *testing.T) {
	b, e := SplitWithKey2("./tmp/ELTbmjn2IZY6EtLFCibQPL4pIyfMvN8jQS67ntPlFaFo3NUkM3PpCFBgMivKk67W", "./split/ELTbmjn2IZY6EtLFCibQPL4pIyfMvN8jQS67ntPlFaFo3NUkM3PpCFBgMivKk67W", "./split/text_keyfile.key")
	t.Log(string(b), e)
}
