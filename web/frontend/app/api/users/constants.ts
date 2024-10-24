import { FormSchemaUser } from "./forms";

export const JWT_SECRET = process.env.AUTH_SECRET;

export const CreateUser = FormSchemaUser.omit({ id: true });
export { FormSchemaUser };
