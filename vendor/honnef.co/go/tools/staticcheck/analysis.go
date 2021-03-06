package staticcheck

import (
	"flag"

	"honnef.co/go/tools/facts"
	"honnef.co/go/tools/internal/passes/buildssa"
	"honnef.co/go/tools/lint/lintutil"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

func newFlagSet() flag.FlagSet {
	fs := flag.NewFlagSet("", flag.PanicOnError)
	fs.Var(lintutil.NewVersionFlag(), "go", "Target Go version")
	return *fs
}

var Analyzers = map[string]*analysis.Analyzer{
	"SA1000": {
		Name:     "SA1000",
		Run:      callChecker(checkRegexpRules),
		Doc:      Docs["SA1000"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1001": {
		Name:     "SA1001",
		Run:      CheckTemplate,
		Doc:      Docs["SA1001"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA1002": {
		Name:     "SA1002",
		Run:      callChecker(checkTimeParseRules),
		Doc:      Docs["SA1002"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1003": {
		Name:     "SA1003",
		Run:      callChecker(checkEncodingBinaryRules),
		Doc:      Docs["SA1003"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1004": {
		Name:     "SA1004",
		Run:      CheckTimeSleepConstant,
		Doc:      Docs["SA1004"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA1005": {
		Name:     "SA1005",
		Run:      CheckExec,
		Doc:      Docs["SA1005"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA1006": {
		Name:     "SA1006",
		Run:      CheckUnsafePrintf,
		Doc:      Docs["SA1006"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA1007": {
		Name:     "SA1007",
		Run:      callChecker(checkURLsRules),
		Doc:      Docs["SA1007"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1008": {
		Name:     "SA1008",
		Run:      CheckCanonicalHeaderKey,
		Doc:      Docs["SA1008"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA1010": {
		Name:     "SA1010",
		Run:      callChecker(checkRegexpFindAllRules),
		Doc:      Docs["SA1010"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1011": {
		Name:     "SA1011",
		Run:      callChecker(checkUTF8CutsetRules),
		Doc:      Docs["SA1011"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1012": {
		Name:     "SA1012",
		Run:      CheckNilContext,
		Doc:      Docs["SA1012"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA1013": {
		Name:     "SA1013",
		Run:      CheckSeeker,
		Doc:      Docs["SA1013"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA1014": {
		Name:     "SA1014",
		Run:      callChecker(checkUnmarshalPointerRules),
		Doc:      Docs["SA1014"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1015": {
		Name:     "SA1015",
		Run:      CheckLeakyTimeTick,
		Doc:      Docs["SA1015"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA1016": {
		Name:     "SA1016",
		Run:      CheckUntrappableSignal,
		Doc:      Docs["SA1016"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA1017": {
		Name:     "SA1017",
		Run:      callChecker(checkUnbufferedSignalChanRules),
		Doc:      Docs["SA1017"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1018": {
		Name:     "SA1018",
		Run:      callChecker(checkStringsReplaceZeroRules),
		Doc:      Docs["SA1018"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1019": {
		Name:     "SA1019",
		Run:      CheckDeprecated,
		Doc:      Docs["SA1019"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer, facts.Deprecated},
		Flags:    newFlagSet(),
	},
	"SA1020": {
		Name:     "SA1020",
		Run:      callChecker(checkListenAddressRules),
		Doc:      Docs["SA1020"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1021": {
		Name:     "SA1021",
		Run:      callChecker(checkBytesEqualIPRules),
		Doc:      Docs["SA1021"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1023": {
		Name:     "SA1023",
		Run:      CheckWriterBufferModified,
		Doc:      Docs["SA1023"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA1024": {
		Name:     "SA1024",
		Run:      callChecker(checkUniqueCutsetRules),
		Doc:      Docs["SA1024"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1025": {
		Name:     "SA1025",
		Run:      CheckTimerResetReturnValue,
		Doc:      Docs["SA1025"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA1026": {
		Name:     "SA1026",
		Run:      callChecker(checkUnsupportedMarshal),
		Doc:      Docs["SA1026"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA1027": {
		Name:     "SA1027",
		Run:      callChecker(checkAtomicAlignment),
		Doc:      Docs["SA1027"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},

	"SA2000": {
		Name:     "SA2000",
		Run:      CheckWaitgroupAdd,
		Doc:      Docs["SA2000"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA2001": {
		Name:     "SA2001",
		Run:      CheckEmptyCriticalSection,
		Doc:      Docs["SA2001"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA2002": {
		Name:     "SA2002",
		Run:      CheckConcurrentTesting,
		Doc:      Docs["SA2002"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA2003": {
		Name:     "SA2003",
		Run:      CheckDeferLock,
		Doc:      Docs["SA2003"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},

	"SA3000": {
		Name:     "SA3000",
		Run:      CheckTestMainExit,
		Doc:      Docs["SA3000"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA3001": {
		Name:     "SA3001",
		Run:      CheckBenchmarkN,
		Doc:      Docs["SA3001"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},

	"SA4000": {
		Name:     "SA4000",
		Run:      CheckLhsRhsIdentical,
		Doc:      Docs["SA4000"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer, facts.TokenFile, facts.Generated},
		Flags:    newFlagSet(),
	},
	"SA4001": {
		Name:     "SA4001",
		Run:      CheckIneffectiveCopy,
		Doc:      Docs["SA4001"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA4002": {
		Name:     "SA4002",
		Run:      CheckDiffSizeComparison,
		Doc:      Docs["SA4002"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA4003": {
		Name:     "SA4003",
		Run:      CheckExtremeComparison,
		Doc:      Docs["SA4003"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA4004": {
		Name:     "SA4004",
		Run:      CheckIneffectiveLoop,
		Doc:      Docs["SA4004"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA4006": {
		Name:     "SA4006",
		Run:      CheckUnreadVariableValues,
		Doc:      Docs["SA4006"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, facts.Generated},
		Flags:    newFlagSet(),
	},
	"SA4008": {
		Name:     "SA4008",
		Run:      CheckLoopCondition,
		Doc:      Docs["SA4008"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA4009": {
		Name:     "SA4009",
		Run:      CheckArgOverwritten,
		Doc:      Docs["SA4009"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA4010": {
		Name:     "SA4010",
		Run:      CheckIneffectiveAppend,
		Doc:      Docs["SA4010"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA4011": {
		Name:     "SA4011",
		Run:      CheckScopedBreak,
		Doc:      Docs["SA4011"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA4012": {
		Name:     "SA4012",
		Run:      CheckNaNComparison,
		Doc:      Docs["SA4012"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA4013": {
		Name:     "SA4013",
		Run:      CheckDoubleNegation,
		Doc:      Docs["SA4013"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA4014": {
		Name:     "SA4014",
		Run:      CheckRepeatedIfElse,
		Doc:      Docs["SA4014"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA4015": {
		Name:     "SA4015",
		Run:      callChecker(checkMathIntRules),
		Doc:      Docs["SA4015"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA4016": {
		Name:     "SA4016",
		Run:      CheckSillyBitwiseOps,
		Doc:      Docs["SA4016"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, facts.TokenFile},
		Flags:    newFlagSet(),
	},
	"SA4017": {
		Name:     "SA4017",
		Run:      CheckPureFunctions,
		Doc:      Docs["SA4017"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, facts.Purity},
		Flags:    newFlagSet(),
	},
	"SA4018": {
		Name:     "SA4018",
		Run:      CheckSelfAssignment,
		Doc:      Docs["SA4018"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer, facts.Generated, facts.TokenFile},
		Flags:    newFlagSet(),
	},
	"SA4019": {
		Name:     "SA4019",
		Run:      CheckDuplicateBuildConstraints,
		Doc:      Docs["SA4019"].String(),
		Requires: []*analysis.Analyzer{facts.Generated},
		Flags:    newFlagSet(),
	},
	"SA4020": {
		Name:     "SA4020",
		Run:      CheckUnreachableTypeCases,
		Doc:      Docs["SA4020"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA4021": {
		Name:     "SA4021",
		Run:      CheckSingleArgAppend,
		Doc:      Docs["SA4021"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer, facts.Generated, facts.TokenFile},
		Flags:    newFlagSet(),
	},

	"SA5000": {
		Name:     "SA5000",
		Run:      CheckNilMaps,
		Doc:      Docs["SA5000"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA5001": {
		Name:     "SA5001",
		Run:      CheckEarlyDefer,
		Doc:      Docs["SA5001"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA5002": {
		Name:     "SA5002",
		Run:      CheckInfiniteEmptyLoop,
		Doc:      Docs["SA5002"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA5003": {
		Name:     "SA5003",
		Run:      CheckDeferInInfiniteLoop,
		Doc:      Docs["SA5003"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA5004": {
		Name:     "SA5004",
		Run:      CheckLoopEmptyDefault,
		Doc:      Docs["SA5004"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA5005": {
		Name:     "SA5005",
		Run:      CheckCyclicFinalizer,
		Doc:      Docs["SA5005"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA5007": {
		Name:     "SA5007",
		Run:      CheckInfiniteRecursion,
		Doc:      Docs["SA5007"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA5008": {
		Name:     "SA5008",
		Run:      CheckStructTags,
		Doc:      Docs["SA5008"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA5009": {
		Name:     "SA5009",
		Run:      callChecker(checkPrintfRules),
		Doc:      Docs["SA5009"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},

	"SA6000": {
		Name:     "SA6000",
		Run:      callChecker(checkRegexpMatchLoopRules),
		Doc:      Docs["SA6000"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA6001": {
		Name:     "SA6001",
		Run:      CheckMapBytesKey,
		Doc:      Docs["SA6001"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA6002": {
		Name:     "SA6002",
		Run:      callChecker(checkSyncPoolValueRules),
		Doc:      Docs["SA6002"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer},
		Flags:    newFlagSet(),
	},
	"SA6003": {
		Name:     "SA6003",
		Run:      CheckRangeStringRunes,
		Doc:      Docs["SA6003"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA6005": {
		Name:     "SA6005",
		Run:      CheckToLowerToUpperComparison,
		Doc:      Docs["SA6005"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},

	"SA9001": {
		Name:     "SA9001",
		Run:      CheckDubiousDeferInChannelRangeLoop,
		Doc:      Docs["SA9001"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA9002": {
		Name:     "SA9002",
		Run:      CheckNonOctalFileMode,
		Doc:      Docs["SA9002"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	"SA9003": {
		Name:     "SA9003",
		Run:      CheckEmptyBranch,
		Doc:      Docs["SA9003"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, facts.TokenFile, facts.Generated},
		Flags:    newFlagSet(),
	},
	"SA9004": {
		Name:     "SA9004",
		Run:      CheckMissingEnumTypesInDeclaration,
		Doc:      Docs["SA9004"].String(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    newFlagSet(),
	},
	// Filtering generated code because it may include empty structs generated from data models.
	"SA9005": {
		Name:     "SA9005",
		Run:      callChecker(checkNoopMarshal),
		Doc:      Docs["SA9005"].String(),
		Requires: []*analysis.Analyzer{buildssa.Analyzer, valueRangesAnalyzer, facts.Generated, facts.TokenFile},
		Flags:    newFlagSet(),
	},
}
