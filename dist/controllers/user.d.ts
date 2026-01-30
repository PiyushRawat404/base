import { Request, Response } from "express";
declare const getUser: (req: Request, res: Response) => Promise<void>;
declare const newUser: (req: Request, res: Response) => Promise<Response<any, Record<string, any>>>;
export { getUser, newUser };
//# sourceMappingURL=user.d.ts.map