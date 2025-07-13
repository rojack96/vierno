import { useNavigate } from "react-router-dom";
import "./Card.css";

interface CardProps {
    elements: string[];
    onClick?: (value: string) => void;
}


const Card = ({ elements, onClick }: CardProps) => {
    const navigate = useNavigate();

    const handleClick = (value: string) => {
        if (onClick) {
            onClick(value); // usa la callback se fornita
        } else {
            navigate(`/folders/${value}`); // fallback
        }
    };

    return (
        <main className="card-container">
            <div className="card-grid">
                {elements.map((value, idx) => (
                    <div key={idx} className="card" onClick={() => handleClick(value)}>
                        <h3>{value}</h3>
                    </div>
                ))}
            </div>
        </main>
    );
};

export default Card;
