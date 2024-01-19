import React from "react";
import { Col, Form, Row } from "react-bootstrap";

import StringSchema from "../../model/schema/StringSchema";
import { IStringSchemaType } from "../../model/schema/type_schema";
import EditorOptionModal from "../node_component/EditorOptionModal";
import EnumField from "../node_component/EnumField";
import GenericField from "../node_component/GenericField";
import HintText from "../node_component/HintText";
import InstructionIcon from "../node_component/InstructionIcon";
import OptionsButtons from "../node_component/OptionsButtons";
import SpaceFront from "../node_component/SpaceFront";
import { Hint, IGenericFieldOptions, IOptionsButtonsAttr } from "../node_component/type_NodeComponent";
import SchemaEditor from "./SchemaEditor";
import { ISchemaEditorProps, ISchemaEditorState, IStringEditorField } from "./type_SchemaEditor";

class StringSchemaEditor extends SchemaEditor<IStringSchemaType, IStringEditorField> {
    protected optionsButtonsAttr: IOptionsButtonsAttr;
    protected genericFieldOptions: IGenericFieldOptions;
    public schema: StringSchema;

    protected optionModalRef: React.RefObject<EditorOptionModal>;
    protected genericFieldRef: React.RefObject<GenericField>;
    private hintTextRef: React.RefObject<HintText>;

    constructor(props: ISchemaEditorProps<IStringSchemaType>) {
        super(props);

        this.optionModalRef = React.createRef<EditorOptionModal>();
        this.genericFieldRef = React.createRef<GenericField>();
        this.hintTextRef = React.createRef<HintText>();

        this.schema = new StringSchema(props.schema, props.field);

        this.optionsButtonsAttr = {
            hasChild: false,
            hasSibling: true,
            isDeleteable: true,
            isOptionExist: true,
            ...props, // override hasSibling, isDeleteable
        };

        this.genericFieldOptions = {
            ...props, // override isRequiredFieldReadonly, isNameFieldReadonly
        };

        this.state = {
            currentField: this.schema.getDefaultField(),
        };
    }

    componentDidMount(): void {
        if (this.state.currentField.minLength > this.state.currentField.maxLength)
            this.hintTextRef.current?.add(Hint.Warn.MIN_GT_MAX_LENGTH);
    }

    componentDidUpdate(prevProps: ISchemaEditorProps<IStringSchemaType>, prevState: ISchemaEditorState<IStringEditorField>): void {
        if (
            // NaN === NaN (get false)
            // NaN !== NaN (get true)
            (prevState.currentField.minLength !== this.state.currentField.minLength &&
                !(isNaN(prevState.currentField.minLength) && isNaN(this.state.currentField.minLength))) ||
            (prevState.currentField.maxLength !== this.state.currentField.maxLength &&
                !(isNaN(prevState.currentField.maxLength) && isNaN(this.state.currentField.maxLength)))
        ) {
            if (this.state.currentField.minLength > this.state.currentField.maxLength) {
                this.hintTextRef.current?.add(Hint.Warn.MIN_GT_MAX_LENGTH);
            } else {
                this.hintTextRef.current?.remove(Hint.Warn.MIN_GT_MAX_LENGTH);
            }
        }
    }

    exportSchema(): IStringSchemaType {
        return this.schema.exportSchema();
    }

    updateEnum(index?: number, changeEvent?: React.ChangeEvent<HTMLInputElement>): void {
        if (index === undefined && changeEvent === undefined) this.schema.addEnum();
        else if (index !== undefined && changeEvent === undefined) this.schema.deleteEnum(index);
        else if (index !== undefined && changeEvent !== undefined) this.schema.updateEnum(index, changeEvent);

        this.setState({ currentField: this.schema.getCurrentField() });
    }

    render(): JSX.Element {
        return (
            <div className="my-1">
                <Row>
                    <SpaceFront depth={this.props.depth} />

                    <Col>
                        <HintText ref={this.hintTextRef} />

                        <Form>
                            <Row>
                                <Col lg={11}>
                                    <GenericField
                                        ref={this.genericFieldRef}
                                        schemaType={this.schema}
                                        options={this.genericFieldOptions}
                                        changeType={this.props.changeType}
                                        changeName={this.props.changeName}
                                    />
                                </Col>
                                <Col lg={1}>
                                    <OptionsButtons
                                        buttonOptions={this.optionsButtonsAttr}
                                        delete={this.delete.bind(this)}
                                        addChild={this.addChild.bind(this)}
                                        addSibling={this.addSibling.bind(this)}
                                        showOptionModal={this.showOptionModal.bind(this)}
                                    />
                                </Col>
                                <EditorOptionModal
                                    clearOptionFieldForm={this.clearOptionField.bind(this)}
                                    resetOptionFiledForm={this.resetOptionField.bind(this)}
                                    ref={this.optionModalRef}
                                >
                                    <Form>
                                        <Form.Group as={Row}>
                                            <Form.Label column lg="2" htmlFor="Default">
                                                Default
                                                <InstructionIcon text="This keyword can be used to supply a default string value associated with this string schema." />
                                            </Form.Label>
                                            <Col lg="10">
                                                <Form.Control
                                                    type="text"
                                                    id="Default"
                                                    value={this.state.currentField.default}
                                                    onChange={this.recordField.bind(this, "default")}
                                                />
                                            </Col>
                                        </Form.Group>

                                        <Form.Group as={Row}>
                                            <Form.Label column lg="2" htmlFor="MinLength">
                                                Min Length
                                                <InstructionIcon text="A string instance would be valid if its length is greate than or equal to this keyword." />
                                            </Form.Label>
                                            <Col lg="4">
                                                <Form.Control
                                                    type="number"
                                                    min="0"
                                                    id="MinLength"
                                                    value={this.state.currentField.minLength}
                                                    onChange={this.recordField.bind(this, "minLength")}
                                                />
                                            </Col>
                                            <Form.Label column lg="2" htmlFor="MaxLength">
                                                Max Length
                                                <InstructionIcon text="A string instance would be valid if its length is less than or equal to this keyword." />
                                            </Form.Label>
                                            <Col lg="4">
                                                <Form.Control
                                                    type="number"
                                                    min="0"
                                                    id="MaxLength"
                                                    value={this.state.currentField.maxLength}
                                                    onChange={this.recordField.bind(this, "maxLength")}
                                                />
                                            </Col>
                                        </Form.Group>

                                        <Form.Group as={Row} controlId="Format">
                                            <Form.Label column lg="2">
                                                Format
                                            </Form.Label>
                                            <Col lg="10">
                                                <Form.Control
                                                    as="select"
                                                    value={this.state.currentField.format}
                                                    onChange={this.recordField.bind(this, "format")}
                                                >
                                                    <option disabled hidden value="">
                                                        {" "}
                                                    </option>
                                                    {[
                                                        "date-time",
                                                        "time",
                                                        "date",
                                                        "email",
                                                        "idn-email",
                                                        "hostname",
                                                        "idn-hostname",
                                                        "ipv4",
                                                        "ipv6",
                                                        "uri",
                                                        "uri-reference",
                                                        "iri",
                                                        "iri-reference",
                                                        "uri-template",
                                                        "json-pointer",
                                                        "relative-json-pointer",
                                                        "regex",
                                                    ].map((v, i) => (
                                                        <option key={i} value={v}>
                                                            {v}
                                                        </option>
                                                    ))}
                                                </Form.Control>
                                            </Col>
                                        </Form.Group>

                                        <Form.Group as={Row} controlId="Pattern">
                                            <Form.Label column lg="2">
                                                Pattern
                                                <InstructionIcon text="The value of this keyword should be a regular expression. A string instance is valid if the regular expression mathes the string instance." />
                                            </Form.Label>
                                            <Col lg="10">
                                                <Form.Control
                                                    type="text"
                                                    placeholder="Regular Expression"
                                                    value={this.state.currentField.pattern}
                                                    onChange={this.recordField.bind(this, "pattern")}
                                                />
                                            </Col>
                                        </Form.Group>

                                        <Form.Group as={Row} controlId="const">
                                            <Form.Label column lg="2">
                                                Constant
                                                <InstructionIcon text="A string instance would be valid if its value is equal to this keyword." />
                                            </Form.Label>
                                            <Col lg="10">
                                                <Form.Control
                                                    type="text"
                                                    placeholder="Restricted Value"
                                                    value={this.state.currentField.const}
                                                    onChange={this.recordField.bind(this, "const")}
                                                />
                                            </Col>
                                        </Form.Group>

                                        <Form.Group as={Row} contolId="enum">
                                            <Col lg="12">
                                                <EnumField
                                                    type="text"
                                                    width={10}
                                                    value={this.state.currentField.enum}
                                                    add={(): void => this.updateEnum()}
                                                    update={this.updateEnum.bind(this)}
                                                    delete={(index: number): void => this.updateEnum(index)}
                                                />
                                            </Col>
                                        </Form.Group>
                                    </Form>
                                </EditorOptionModal>
                            </Row>
                        </Form>
                    </Col>
                </Row>
            </div>
        );
    }
}

export default StringSchemaEditor;
