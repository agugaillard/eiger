package rdiff_test

import (
	"testing"

	"github.com/agugaillard/eiger/rdiff"
	"github.com/stretchr/testify/assert"
)

func Test_Patch(t *testing.T) {
	rdiff := rdiff.NewDefaultRdiff(logger)
	loremImpsum := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla nisl enim, consectetur quis quam consequat, pharetra tempus enim. Fusce iaculis libero vitae ipsum accumsan efficitur. Fusce iaculis est et justo sollicitudin, sed porttitor augue sagittis. Mauris aliquam nisl nibh, sed tempus magna venenatis ac. Curabitur molestie nisl elit, suscipit egestas ex aliquam ac. Donec dignissim, mauris nec malesuada pellentesque, ipsum sem porttitor est, quis laoreet urna orci a leo. Cras tincidunt porttitor sapien, quis cursus metus pulvinar id. Pellentesque nec mollis eros. Fusce sagittis vehicula ligula, nec ullamcorper sapien sagittis non.")
	modifiedLoremImpsum := []byte("Proin finibus ullamcorper ante sit amet egestas. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla nisl enim, consectetur quis quam consequat, pharetra tempus enim. Fusce iaculis libero vitae ipsum accumsan efficitur. Fusce iaculis est et justo sollicitudin, sed porttitor augue sagittis. Mauris aliquam nisl nibh, sed tempus magna venenatis ac. Nulla ex metus, malesuada eget ultricies vel, fermentum quis nisl. Etiam ac venenatis tellus. Curabitur molestie nisl elit, suscipit egestas ex aliquam ac. Donec dignissim, mauris nec malesuada pellentesque, ipsum sem porttitor est, quis laoreet urna orci a leo. Cras tincidunt porttitor sapien, quis cursus metus pulvinar id. Pellentesque nec mollis eros.")

	chunkSizes := []uint{2, 3, 5, 10, 32, 64}
	for _, chunkSize := range chunkSizes {
		signature := rdiff.Signature(loremImpsum, chunkSize)
		delta := rdiff.Delta(modifiedLoremImpsum, signature)
		output := rdiff.Patch(loremImpsum, delta)
		assert.Equal(t, string(modifiedLoremImpsum), output)
	}
}
