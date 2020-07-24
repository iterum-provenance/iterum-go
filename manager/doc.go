// Package manager contains interactive features for the Iterum manager.
// It checks the relevant environment variables and contains a goroutine called
// UpstreamChecker which is responsible for checking whether previous transformations have concluded.
// This is the main routine that enables Iterum's automated completion.
package manager
