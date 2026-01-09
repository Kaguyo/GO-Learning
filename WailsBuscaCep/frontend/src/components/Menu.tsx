import React, { useEffect, useState } from 'react';
import "./Menu.css";

interface MenuItem {
  id: string;
  label: string;
  icon?: string;
}

type MenuItemKey = "atualizacao-automatica" | "buscar-localizacao" | "inserir-localizacao";

function Menu(): JSX.Element {
  const [focusedServices, setFocusedServices] = useState<MenuItemKey[]>([]);

  const toggleServiceFocus = (key: MenuItemKey) => {
    setFocusedServices((prev) => {
      if (prev.includes(key)) {
        return prev.filter(item => item !== key);
      }
      return [...prev, key];
    });
  };

  const serviceItemsMap: Record<MenuItemKey, MenuItem> = {
    "atualizacao-automatica": { id: '1', label: 'Atualização Automática' },
    "buscar-localizacao": { id: '2', label: 'Buscar Localização' },
    "inserir-localizacao": { id: '3', label: 'Inserir Localização' },
  };

  // Transformamos o mapa em array para renderizar, verificando o estado de foco dinamicamente
  const menuKeys = Object.keys(serviceItemsMap) as MenuItemKey[];

  useEffect(() => {
    console.log("Serviços ativos:", focusedServices);
  }, [focusedServices]);

  return (
    <aside className="menu-container" style={{ width: '50%', border: '1px solid #ddd' }}>
      <nav className="menu-navbar">
        <button className="nav-button" onClick={() => console.log('Foco: Serviços')}>
            Serviços
        </button>
        <button className="nav-button" onClick={() => console.log('Foco: Configurações')}>
            Configurações
        </button>
      </nav>

      <main className="menu-actions" style={{ padding: '15px' }}>
        <h3 style={{ fontSize: '0.8rem', color: '#666', textTransform: 'uppercase', }}>Serviços</h3>
        <ul style={{ listStyle: 'none', padding: 0, paddingBottom: "30px"}}>
          {menuKeys.map((key) => {
            const item = serviceItemsMap[key];
            const isFocused = focusedServices.includes(key);

            return (
            <li key={item.id} className="menu-item-wrapper" style={{ paddingBottom: "20px", width: "75%", marginLeft:"50%", transform: "translate(-50%)"}}>
                <button
                onClick={() => toggleServiceFocus(key)}
                className={`menu-button ${isFocused ? 'focused' : ''}`}
                >
                <span>{item.label}</span>
                {!isFocused && <span className="arrow-icon">▼</span>}
                </button>

                <div className={`expandable-grid ${isFocused ? 'expanded' : ''}`}>
                    <div className="expandable-container">
                        <div className="inner-content" style={{color: "black"}}>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                            <h3>{item.label}</h3>
                        </div>
                    </div>
                </div>
            </li>
            );
          })}
        </ul>
      </main>
    </aside>
  );
}

export default Menu;