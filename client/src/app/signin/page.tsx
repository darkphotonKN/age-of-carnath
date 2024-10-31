"use client";
import { Button } from "@/components/Button";
import axios from "axios";
import { ErrorMessage, Field, Form, Formik } from "formik";
import { useRouter } from "next/navigation";
import * as Yup from "yup";

const validationSchema = Yup.object({
  email: Yup.string().email().required("Email is required"),
  password: Yup.string().min(6).required("Password is required"),
});

function SignInPage() {
  const router = useRouter();

  return (
    <div className="h-[400px] w-screen flex justify-center items-center">
      <Formik
        initialValues={{
          name: "",
          phone: "",
          birthday: "",
        }}
        validationSchema={validationSchema}
        onSubmit={async (values) => {
          try {
            const response = await axios.post("url", { ...values });
            console.log("Success:", response.data);

            // handle token storage and redirect to home
            router.push("/");
          } catch (error) {
            console.error("Error:", error);
            // handle error (e.g., show a message or redirect)
          }
        }}
      >
        {({ isSubmitting }) => (
          <Form className="flex flex-col gap-3 text-lg">
            <div className="text-lg font-semibold">SIGN IN</div>
            <div>
              <label className="mr-2" htmlFor="email">
                Email
              </label>
              <Field
                type="email"
                name="email"
                className="rounded-sm py-[2px] px-[8px]"
              />
              <ErrorMessage
                className="text-red-500 mt-2"
                name="email"
                component="div"
              />
            </div>

            <div>
              <label className="mr-2" htmlFor="password">
                Password
              </label>
              <Field
                type="password"
                name="password"
                className="rounded-sm py-[2px] px-[8px]"
              />
              <ErrorMessage
                className="text-red-500 mt-2"
                name="password"
                component="div"
              />
            </div>

            <Button
              className="mt-3"
              type="submit"
              variant="default"
              disabled={isSubmitting}
            >
              Submit
            </Button>
          </Form>
        )}
      </Formik>
    </div>
  );
}

export default SignInPage;
