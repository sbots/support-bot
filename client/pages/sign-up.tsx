import React from "react";
import { useForm } from "react-hook-form";
import { Input } from "../components/Input/Input";
import Link from "next/link";

type Inputs = {
  email: string;
  password: string;
  firstName: string;
  lastName: string;
};

const SignUp = () => {
  const { register, handleSubmit } = useForm<Inputs>();

  const onSubmit = (data: Inputs) => {
    console.log(data);
  };
  return (
    <div className="min-h-screen bg-white flex">
      <div className="hidden lg:block relative w-0 flex-1">
        <img
          className="absolute inset-0 h-full w-full object-cover"
          src="https://images.unsplash.com/photo-1514924013411-cbf25faa35bb?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=936&q=80"
          alt=""
        />
      </div>
      <div className="flex-1 flex flex-col justify-center py-12 px-4 sm:px-6 lg:flex-none lg:px-20 xl:px-24">
        <div className="mx-auto w-full max-w-sm lg:w-96">
          <div>
            <img
              className="h-12 w-auto"
              src="https://tailwindui.com/img/logos/workflow-mark-indigo-600.svg"
              alt="Workflow"
            />
            <h2 className="mt-6 text-3xl font-extrabold text-gray-900">
              Create your account
            </h2>
          </div>

          <div className="mt-8">
            <div className="mt-6">
              <form className="space-y-6" onSubmit={handleSubmit(onSubmit)}>
                <Input
                  label="Email address"
                  type="text"
                  {...register("email", {
                    required: "Email address not valid",
                    pattern: /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/i,
                  })}
                />
                <Input
                  label="Password"
                  type="text"
                  className="space-y-1"
                  {...register("password", {
                    required: true,
                  })}
                />
                <Input
                  label="First Name"
                  type="text"
                  className="space-y-1"
                  {...register("firstName", {
                    required: true,
                  })}
                />
                <Input
                  label="Last Name"
                  type="text"
                  className="space-y-1"
                  {...register("lastName", {
                    required: true,
                  })}
                />

                <div className="flex items-center justify-between">
                  <div className="text-sm">
                    <Link href="/sign-in">
                      <a className="font-medium text-indigo-600 hover:text-indigo-500">
                        Have an account? sign-in
                      </a>
                    </Link>
                  </div>
                </div>

                <div>
                  <button
                    type="submit"
                    className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                  >
                    Sign up
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
export default SignUp;
