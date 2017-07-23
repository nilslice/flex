package flex

import "testing"

func TestDont_cache_computed_flex_basis_between_layouts(t *testing.T) {
	config := YGConfigNew()
	YGConfigSetExperimentalFeatureEnabled(config, YGExperimentalFeatureWebFlexBasis, true)

	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetHeightPercent(root, 100)
	YGNodeStyleSetWidthPercent(root, 100)

	rootChild0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetFlexBasisPercent(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, 100, YGUndefined, DirectionLTR)
	YGNodeCalculateLayout(root, 100, 100, DirectionLTR)

	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

}

func TestRecalculate_resolvedDimonsion_onchange(t *testing.T) {
	root := YGNodeNew()

	rootChild0 := YGNodeNew()
	YGNodeStyleSetMinHeight(rootChild0, 10)
	YGNodeStyleSetMaxHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	YGNodeStyleSetMinHeight(rootChild0, YGUndefined)
	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

}
