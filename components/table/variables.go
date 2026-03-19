package table

type TableModifier string

func (tm TableModifier) String() string {
	return string(tm)
}

const (
	Divider TableModifier = "uk-table-divider"
	Striped TableModifier = "uk-table-striped"
	Hover   TableModifier = "uk-table-hover"

	Justify            TableModifier = "uk-table-justify" // To remove the outer padding of the first and last columns so that they are flush with the table
	AlignMiddle        TableModifier = "uk-table-middle"  // To vertically center table content, just add the .uk-table-middle class. You can also apply the class to <tr> or <td> elements for a more specific selection.
	ResponsiveOverflow TableModifier = "uk-overflow-auto" // If your table happens to be wider than its container element or would eventually get too big on a specific viewport width, just wrap it inside a <div> element and add the .uk-overflow-auto class. This creates a container that provides a horizontal scrollbar whenever the elements inside it are wider than the container itself.
	ResponsiveStack    TableModifier = "uk-table-responsive"
	ColumnWidthShrink  TableModifier = "uk-table-shrink" // Add this class to a <th> or <td> element to reduce the column width to fit its content.
	ColumnWidthExpand  TableModifier = "uk-table-expand" // Add this class to a <th> or <td> element to expand the column width to fill the remaining space and apply a min-width.
	// ColumnWidth... TableModifier = "w-*" Add one of the width utility classes from Tailwind CSS to a <th> or <td> element to modify the column width., e.g. w-40
	//Utilities
	LinkEntireCell TableModifier = "uk-table-link"     // To link an entire cell, add this class to a <th>or <td> element and insert an <a> element. Add the .uk-link-resetclass from the Link component to reset the default link styling.
	PreserveWidth  TableModifier = "uk-preserve-width" // Since images are responsive by default in UIkit, using an image inside a table cell with the .uk-table-shrink class would reduce the image width to 0. To prevent this behavior, add the .uk-preserve-width class to the <img> element.
	TruncateText   TableModifier = "uk-text-truncate"  // When applying a fixed width to a table cell, you might want to add this class to the <td> element to truncate the text.
	NoWrap         TableModifier = "text-nowrap"       // Add this class from the Tailwind CSS to keep text from wrapping, for example when using the .uk-table-shrink class.
)

type TableSize string

func (t TableSize) String() string {
	return string(t)
}

const (
	SizeSm TableSize = "uk-table-sm"
	SizeMd TableSize = "uk-table-md"
	SizeLg TableSize = "uk-table-lg"
)
