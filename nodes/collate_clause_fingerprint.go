// Auto-generated - DO NOT EDIT

package pg_query

func (node CollateClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CollateClause")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx, "Arg")
	}

	node.Collname.Fingerprint(ctx, "Collname")
	// Intentionally ignoring node.Location for fingerprinting
}
