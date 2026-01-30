import mongoose from "mongoose";
export declare const User: mongoose.Model<{
    clientNo: number;
    fullName: string;
    phoneNo: number;
    gender: string;
    problem: string;
}, {}, {}, {
    id: string;
}, mongoose.Document<unknown, {}, {
    clientNo: number;
    fullName: string;
    phoneNo: number;
    gender: string;
    problem: string;
}, {
    id: string;
}, mongoose.DefaultSchemaOptions> & Omit<{
    clientNo: number;
    fullName: string;
    phoneNo: number;
    gender: string;
    problem: string;
} & {
    _id: mongoose.Types.ObjectId;
} & {
    __v: number;
}, "id"> & {
    id: string;
}, mongoose.Schema<any, mongoose.Model<any, any, any, any, any, any, any>, {}, {}, {}, {}, mongoose.DefaultSchemaOptions, {
    clientNo: number;
    fullName: string;
    phoneNo: number;
    gender: string;
    problem: string;
}, mongoose.Document<unknown, {}, {
    clientNo: number;
    fullName: string;
    phoneNo: number;
    gender: string;
    problem: string;
}, {
    id: string;
}, mongoose.ResolveSchemaOptions<mongoose.DefaultSchemaOptions>> & Omit<{
    clientNo: number;
    fullName: string;
    phoneNo: number;
    gender: string;
    problem: string;
} & {
    _id: mongoose.Types.ObjectId;
} & {
    __v: number;
}, "id"> & {
    id: string;
}, {
    [path: string]: mongoose.SchemaDefinitionProperty<undefined, any, any>;
} | {
    [x: string]: mongoose.SchemaDefinitionProperty<any, any, mongoose.Document<unknown, {}, {
        clientNo: number;
        fullName: string;
        phoneNo: number;
        gender: string;
        problem: string;
    }, {
        id: string;
    }, mongoose.ResolveSchemaOptions<mongoose.DefaultSchemaOptions>> & Omit<{
        clientNo: number;
        fullName: string;
        phoneNo: number;
        gender: string;
        problem: string;
    } & {
        _id: mongoose.Types.ObjectId;
    } & {
        __v: number;
    }, "id"> & {
        id: string;
    }> | undefined;
}, {
    clientNo: number;
    fullName: string;
    phoneNo: number;
    gender: string;
    problem: string;
} & {
    _id: mongoose.Types.ObjectId;
} & {
    __v: number;
}>, {
    clientNo: number;
    fullName: string;
    phoneNo: number;
    gender: string;
    problem: string;
} & {
    _id: mongoose.Types.ObjectId;
} & {
    __v: number;
}>;
//# sourceMappingURL=user.d.ts.map