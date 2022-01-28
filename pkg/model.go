package pkg

//xcode project.pbxproj

type Project struct {

	//rootObject
	RootObject string `json:"rootObject"`
	//objectVersion = 48;
	ObjectVersion string `json:"{"`

	//archiveVersion = 1;
	ArchiveVersion string `json:"archiveVersion"`

	//classes
	Classes []string `json:"classes"`

	//objects
	Objects map[string]Object `json:"objects"`
}

type Object interface {
	getIsa() string
}

type Sections struct {

	// /* Begin PBXBuildFile section */
	PBXBuildFile []PBXBuildFileItem
	//  /* Begin PBXFileReference section */

	PBXFileReference []PBXFileReferenceItem

	//PBXFrameworksBuildPhase
	PBXFrameworksBuildPhase []PBXFrameworksBuildPhaseItem

	/* Begin PBXGroup section */
	PBXGroup []PBXGroupItem
}

type BaseItem struct {
	UUID string
	Isa  string `json:"isa"`
}

func (item *BaseItem) getIsa() string {
	return item.Isa
}

type PBXBuildFileItem struct {
	// C8540F9E20616EA100F75258
	UUID string
	// /* ViewController.swift in Sources */
	Comment string
	// isa = PBXBuildFile;
	ISA string
	//fileRef = C8540F9D20616EA100F75258 /* ViewController.swift */;
	FileRef string
}

type PBXFileReferenceItem struct {
	Comment string
	UUID    string
	// isa = PBXFileReference;
	ISA string
	//includeInIndex = 1;
	includeInIndex int
	// lastKnownFileType = text.xcconfig;
	// lastKnownFileType = file.storyboard;
	lastKnownFileType string
	//name = "Pods-MultiPeer_Sample.release.xcconfig";
	name string
	//path = "Pods/Target Support Files/Pods-MultiPeer_Sample/Pods-MultiPeer_Sample.release.xcconfig";
	path string
	//sourceTree = "<group>";
	sourceTree string
}

//
//C8540F9520616EA100F75258 /* Frameworks */ = {
//isa = PBXFrameworksBuildPhase;
//buildActionMask = 2147483647;
//files = (
//DCB05CA8CB956B53F258C91D /* Pods_MultiPeer_Sample.framework in Frameworks */,
//);
//runOnlyForDeploymentPostprocessing = 0;
//};
type PBXFrameworksBuildPhaseItem struct {
	// C8540F9520616EA100F75258
	UUID string
	//isa = PBXFrameworksBuildPhase;
	ISA string
	// buildActionMask = 2147483647;
	buildActionMask string
	//DCB05CA8CB956B53F258C91D /* Pods_MultiPeer_Sample.framework in Frameworks */,
	files []string
	// runOnlyForDeploymentPostprocessing = 0;
	runOnlyForDeploymentPostprocessing int
}

type PBXGroupItem struct {

	// C8540F9520616EA100F75258
	UUID string
	//isa = PBXFrameworksBuildPhase;
	ISA string
	//name = Frameworks;
	name string
	//sourceTree = "<group>";
	sourceTree string
	children   []string
}

type PBXNativeTargetItem struct {
	//PBXNativeTarget
	UUID string
	//isa = PBXFrameworksBuildPhase;
	ISA string
	// /* MultiPeer_Sample */
	comment string
}
