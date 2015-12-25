// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SecLabelStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SecLabelStmt")

	if node.Label != nil {
		ctx.WriteString(*node.Label)
	}

	node.Objargs.Fingerprint(ctx, "Objargs")
	node.Objname.Fingerprint(ctx, "Objname")
	ctx.WriteString(strconv.Itoa(int(node.Objtype)))

	if node.Provider != nil {
		ctx.WriteString(*node.Provider)
	}
}
