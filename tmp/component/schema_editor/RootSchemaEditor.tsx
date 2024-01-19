import React from "react";

import { ISchemaType } from "../../model/schema/type_schema";
import { DataType } from "../../type";
import SchemaEditorFactory from "./SchemaEditorFactory";
import { ISchemaEditorType } from "./type_SchemaEditor";

interface RootSchemaEditorProps {
    schema?: ISchemaType;
}

class RootSchemaEditor extends React.Component<RootSchemaEditorProps, { type: DataType }> {
    private editorRef: React.RefObject<ISchemaEditorType>;

    constructor(props: RootSchemaEditorProps) {
        super(props);

        this.editorRef = React.createRef<ISchemaEditorType>();

        if (props.schema) {
            this.state = { type: props.schema.type };
        } else {
            this.state = { type: DataType.Object };
        }
    }

    changeType(type: DataType): void {
        this.setState({ type });
    }

    changeName(): void {
        console.log(`[Root Node] Change Name!!`);
    }

    exportSchema(): ISchemaType {
        // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
        return this.editorRef.current!.exportSchema();
    }

    render(): JSX.Element {
        return (
            <SchemaEditorFactory
                ref={this.editorRef}
                type={this.state.type}
                schema={this.props.schema}
                depth={0}
                field={{ name: "root", required: true }}
                hasSibling={false}
                isDeleteable={false}
                isRequiredFieldReadonly={true}
                isNameFieldReadonly={true}
                changeType={this.changeType.bind(this)}
                changeName={this.changeName.bind(this)}
            />
        );
    }
}

export default RootSchemaEditor;
