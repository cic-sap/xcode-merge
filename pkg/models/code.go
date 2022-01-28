package models

type BaseItem struct {
	UUID string `json:"-"`
	Isa  string `json:"isa"`
}

func (item BaseItem) getIsa() string {
	return item.Isa
}

type PBXNativeTarget struct {
	BaseItem
	BuildConfigurationList string        `json:"buildConfigurationList"`
	BuildPhases            []string      `json:"buildPhases"`
	BuildRules             []interface{} `json:"buildRules"`
	ProductType            string        `json:"productType"`
	Dependencies           []interface{} `json:"dependencies"`
	Name                   string        `json:"name"`
	ProductName            string        `json:"productName"`
	ProductReference       string        `json:"productReference"`
}
type PBXProject struct {
	BaseItem
	DevelopmentRegion      string   `json:"developmentRegion"`
	HasScannedForEncodings float64  `json:"hasScannedForEncodings"`
	Targets                []string `json:"targets"`
	CompatibilityVersion   string   `json:"compatibilityVersion"`
	KnownRegions           []string `json:"knownRegions"`
	MainGroup              string   `json:"mainGroup"`
	ProductRefGroup        string   `json:"productRefGroup"`
	ProjectDirPath         string   `json:"projectDirPath"`
	ProjectRoot            string   `json:"projectRoot"`
	Attributes             struct {
		LastSwiftUpdateCheck float64 `json:"LastSwiftUpdateCheck"`
		LastUpgradeCheck     float64 `json:"LastUpgradeCheck"`
		TargetAttributes     struct {
			B286128C2745256800E1261A struct {
				CreatedOnToolsVersion float64 `json:"CreatedOnToolsVersion"`
			} `json:"B286128C2745256800E1261A"`
		} `json:"TargetAttributes"`
		BuildIndependentTargetsInParallel float64 `json:"BuildIndependentTargetsInParallel"`
	} `json:"attributes"`
	BuildConfigurationList string `json:"buildConfigurationList"`
}
type PBXBuildFile struct {
	BaseItem
	FileRef string `json:"fileRef"`
}
type PBXFileReference struct {
	BaseItem
	LastKnownFileType string `json:"lastKnownFileType"`
	Path              string `json:"path"`
	SourceTree        string `json:"sourceTree"`
}
type PBXGroup struct {
	BaseItem
	Children   []string `json:"children"`
	SourceTree string   `json:"sourceTree"`
}
type XCBuildConfiguration struct {
	BaseItem
	BuildSettings struct {
		CLANG_WARN_ENUM_CONVERSION                    string  `json:"CLANG_WARN_ENUM_CONVERSION"`
		CLANG_WARN_OBJC_IMPLICIT_RETAIN_SELF          string  `json:"CLANG_WARN_OBJC_IMPLICIT_RETAIN_SELF"`
		SWIFT_COMPILATION_MODE                        string  `json:"SWIFT_COMPILATION_MODE"`
		CLANG_ENABLE_OBJC_ARC                         string  `json:"CLANG_ENABLE_OBJC_ARC"`
		CLANG_WARN_DOCUMENTATION_COMMENTS             string  `json:"CLANG_WARN_DOCUMENTATION_COMMENTS"`
		CLANG_WARN_RANGE_LOOP_ANALYSIS                string  `json:"CLANG_WARN_RANGE_LOOP_ANALYSIS"`
		ENABLE_STRICT_OBJC_MSGSEND                    string  `json:"ENABLE_STRICT_OBJC_MSGSEND"`
		GCC_C_LANGUAGE_STANDARD                       string  `json:"GCC_C_LANGUAGE_STANDARD"`
		GCC_WARN_64_TO_32_BIT_CONVERSION              string  `json:"GCC_WARN_64_TO_32_BIT_CONVERSION"`
		CLANG_CXX_LANGUAGE_STANDARD                   string  `json:"CLANG_CXX_LANGUAGE_STANDARD"`
		CLANG_ENABLE_OBJC_WEAK                        string  `json:"CLANG_ENABLE_OBJC_WEAK"`
		CLANG_WARN_INFINITE_RECURSION                 string  `json:"CLANG_WARN_INFINITE_RECURSION"`
		CLANG_WARN_INT_CONVERSION                     string  `json:"CLANG_WARN_INT_CONVERSION"`
		CLANG_WARN_UNGUARDED_AVAILABILITY             string  `json:"CLANG_WARN_UNGUARDED_AVAILABILITY"`
		CLANG_WARN_UNREACHABLE_CODE                   string  `json:"CLANG_WARN_UNREACHABLE_CODE"`
		ENABLE_NS_ASSERTIONS                          string  `json:"ENABLE_NS_ASSERTIONS"`
		IPHONEOS_DEPLOYMENT_TARGET                    float64 `json:"IPHONEOS_DEPLOYMENT_TARGET"`
		CLANG_WARN_BOOL_CONVERSION                    string  `json:"CLANG_WARN_BOOL_CONVERSION"`
		CLANG_WARN_DIRECT_OBJC_ISA_USAGE              string  `json:"CLANG_WARN_DIRECT_OBJC_ISA_USAGE"`
		MTL_FAST_MATH                                 string  `json:"MTL_FAST_MATH"`
		CLANG_WARN_OBJC_LITERAL_CONVERSION            string  `json:"CLANG_WARN_OBJC_LITERAL_CONVERSION"`
		CLANG_WARN_OBJC_ROOT_CLASS                    string  `json:"CLANG_WARN_OBJC_ROOT_CLASS"`
		GCC_WARN_UNINITIALIZED_AUTOS                  string  `json:"GCC_WARN_UNINITIALIZED_AUTOS"`
		SDKROOT                                       string  `json:"SDKROOT"`
		SWIFT_OPTIMIZATION_LEVEL                      string  `json:"SWIFT_OPTIMIZATION_LEVEL"`
		ALWAYS_SEARCH_USER_PATHS                      string  `json:"ALWAYS_SEARCH_USER_PATHS"`
		CLANG_WARN_NON_LITERAL_NULL_CONVERSION        string  `json:"CLANG_WARN_NON_LITERAL_NULL_CONVERSION"`
		CLANG_WARN_SUSPICIOUS_MOVE                    string  `json:"CLANG_WARN_SUSPICIOUS_MOVE"`
		DEBUG_INFORMATION_FORMAT                      string  `json:"DEBUG_INFORMATION_FORMAT"`
		GCC_WARN_UNUSED_VARIABLE                      string  `json:"GCC_WARN_UNUSED_VARIABLE"`
		MTL_ENABLE_DEBUG_INFO                         string  `json:"MTL_ENABLE_DEBUG_INFO"`
		VALIDATE_PRODUCT                              string  `json:"VALIDATE_PRODUCT"`
		CLANG_CXX_LIBRARY                             string  `json:"CLANG_CXX_LIBRARY"`
		CLANG_WARN_CONSTANT_CONVERSION                string  `json:"CLANG_WARN_CONSTANT_CONVERSION"`
		CLANG_WARN__DUPLICATE_METHOD_MATCH            string  `json:"CLANG_WARN__DUPLICATE_METHOD_MATCH"`
		COPY_PHASE_STRIP                              string  `json:"COPY_PHASE_STRIP"`
		GCC_NO_COMMON_BLOCKS                          string  `json:"GCC_NO_COMMON_BLOCKS"`
		GCC_WARN_UNDECLARED_SELECTOR                  string  `json:"GCC_WARN_UNDECLARED_SELECTOR"`
		GCC_WARN_UNUSED_FUNCTION                      string  `json:"GCC_WARN_UNUSED_FUNCTION"`
		CLANG_ANALYZER_NUMBER_OBJECT_CONVERSION       string  `json:"CLANG_ANALYZER_NUMBER_OBJECT_CONVERSION"`
		CLANG_WARN_DEPRECATED_OBJC_IMPLEMENTATIONS    string  `json:"CLANG_WARN_DEPRECATED_OBJC_IMPLEMENTATIONS"`
		CLANG_WARN_EMPTY_BODY                         string  `json:"CLANG_WARN_EMPTY_BODY"`
		CLANG_WARN_STRICT_PROTOTYPES                  string  `json:"CLANG_WARN_STRICT_PROTOTYPES"`
		GCC_WARN_ABOUT_RETURN_TYPE                    string  `json:"GCC_WARN_ABOUT_RETURN_TYPE"`
		CLANG_ENABLE_MODULES                          string  `json:"CLANG_ENABLE_MODULES"`
		CLANG_WARN_COMMA                              string  `json:"CLANG_WARN_COMMA"`
		CLANG_WARN_QUOTED_INCLUDE_IN_FRAMEWORK_HEADER string  `json:"CLANG_WARN_QUOTED_INCLUDE_IN_FRAMEWORK_HEADER"`
		CLANG_ANALYZER_NONNULL                        string  `json:"CLANG_ANALYZER_NONNULL"`
		CLANG_WARN_BLOCK_CAPTURE_AUTORELEASING        string  `json:"CLANG_WARN_BLOCK_CAPTURE_AUTORELEASING"`
	} `json:"buildSettings"`
	Name string `json:"name"`
}
type XCConfigurationList struct {
	BaseItem
	BuildConfigurations           []string `json:"buildConfigurations"`
	DefaultConfigurationIsVisible float64  `json:"defaultConfigurationIsVisible"`
	DefaultConfigurationName      string   `json:"defaultConfigurationName"`
}
type PBXSourcesBuildPhase struct {
	BaseItem
	BuildActionMask                    float64  `json:"buildActionMask"`
	Files                              []string `json:"files"`
	RunOnlyForDeploymentPostprocessing float64  `json:"runOnlyForDeploymentPostprocessing"`
}
type PBXFrameworksBuildPhase struct {
	BaseItem
	RunOnlyForDeploymentPostprocessing float64       `json:"runOnlyForDeploymentPostprocessing"`
	BuildActionMask                    float64       `json:"buildActionMask"`
	Files                              []interface{} `json:"files"`
}
type PBXResourcesBuildPhase struct {
	BaseItem
	BuildActionMask                    float64  `json:"buildActionMask"`
	Files                              []string `json:"files"`
	RunOnlyForDeploymentPostprocessing float64  `json:"runOnlyForDeploymentPostprocessing"`
}
