const Button = ({
  children,
  onClick,
  selected,
}: {
  children: React.ReactNode;
  onClick: () => void;
  selected: boolean;
}) => {
  return (
    <button
      onClick={onClick}
      className={`bg-blue-500 text-sm text-white px-4 py-2 rounded-md hover:bg-blue-600 transition-colors duration-300 cursor-pointer ${
        selected ? "bg-blue-600" : ""
      }`}
    >
      {children}
    </button>
  );
};

export default Button;
