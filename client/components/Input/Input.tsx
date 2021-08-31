import React, { InputHTMLAttributes } from "react";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label: string;
  type: string;
  classNameBlock?: string;
  name: string;
}

export const Input = React.forwardRef<HTMLInputElement, InputProps>(
  (props: InputProps, ref) => {
    const { label, type, name, classNameBlock, ...rest } = props;
    return (
      <div className={classNameBlock}>
        <label
          htmlFor={type}
          className="block text-sm font-medium text-gray-700"
        >
          {label}
        </label>
        <div className="mt-1">
          <input
            {...rest}
            ref={ref}
            id={type}
            name={name}
            type={type}
            autoComplete={type}
            className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
        </div>
      </div>
    );
  }
);
