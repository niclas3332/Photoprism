package photoprism

import (
	"github.com/photoprism/photoprism/internal/classify"
	"github.com/photoprism/photoprism/internal/nsfw"
	"github.com/photoprism/photoprism/internal/thumb"
	"github.com/photoprism/photoprism/pkg/sanitize"
)

// NSFW returns true if media file might be offensive and detection is enabled.
func (ind *Index) NSFW(jpeg *MediaFile) bool {
	filename, err := jpeg.Thumbnail(Config().ThumbPath(), thumb.Fit720)

	if err != nil {
		log.Error(err)
		return false
	}

	if nsfwLabels, err := ind.nsfwDetector.File(filename); err != nil {
		log.Error(err)
		return false
	} else {
		if nsfwLabels.NSFW(nsfw.ThresholdHigh) {

			log.Warnf("index: %s might contain offensive content", sanitize.Log(jpeg.RelName(Config().OriginalsPath())))
			return true
		}
	}

	return false
}

func (ind *Index) NSFWLabels(jpeg *MediaFile) classify.Labels {
	filename, err := jpeg.Thumbnail(Config().ThumbPath(), thumb.Fit720)

	if err != nil {
		log.Error(err)
		return classify.Labels{}
	}

	if nsfwLabels, err := ind.nsfwDetector.File(filename); err != nil {
		log.Error(err)
		return classify.Labels{}
	} else {
		if nsfwLabels.NSFW(nsfw.ThresholdHigh) {

			labels := classify.Labels{}

			if nsfwLabels.Sexy > nsfw.ThresholdHigh {

				labels = append(labels, classify.Label{

					Name:        "sexy",
					Source:      "nsfw",
					Uncertainty: int(nsfwLabels.Sexy),
					Priority:    int(nsfwLabels.Sexy),
					Categories:  nil,
				})

			}
			if nsfwLabels.Porn > nsfw.ThresholdHigh {

				labels = append(labels, classify.Label{

					Name:        "porn",
					Source:      "nsfw",
					Uncertainty: int(nsfwLabels.Sexy),
					Priority:    int(nsfwLabels.Sexy),
					Categories:  nil,
				})

			}

			if nsfwLabels.Hentai > nsfw.ThresholdHigh {

				labels = append(labels, classify.Label{

					Name:        "hentai",
					Source:      "nsfw",
					Uncertainty: int(nsfwLabels.Sexy),
					Priority:    int(nsfwLabels.Sexy),
					Categories:  nil,
				})

			}

			return labels

		}
	}

	return classify.Labels{}
}
