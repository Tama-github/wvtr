package data

// Weapons

func NewInventory(c []*CurrencyOwned) *Inventory {
	return &Inventory{
		Weapons:    make([]*Weapon, 0),
		Armors:     make([]*Armor, 0),
		Omamoris:   make([]*Omamori, 0),
		Currencies: c,
	}
}

func (i *Inventory) AddWeapon(w *Weapon) {
	i.Weapons = append(i.Weapons, w)
}

func (inv *Inventory) FindWeapon(s *Weapon) int {
	for i, stbl := range inv.Weapons {
		if stbl.ID == s.ID {
			return i
		}
	}
	return -1
}

func (inv *Inventory) RemoveWeapon(s *Weapon) *Weapon {
	i := inv.FindWeapon(s)
	if i >= 0 {
		res := inv.Weapons[i]
		inv.Weapons = append(inv.Weapons[0:i], inv.Weapons[i+1:len(inv.Weapons)]...)
		s.InventoryID = 0
		return res
	}
	return nil
}

// Armors

func (i *Inventory) AddArmor(a *Armor) {
	i.Armors = append(i.Armors, a)
}

func (inv *Inventory) FindArmor(s *Armor) int {
	for i, stbl := range inv.Armors {
		if stbl.ID == s.ID {
			return i
		}
	}
	return -1
}

func (inv *Inventory) RemoveArmor(s *Armor) *Armor {
	i := inv.FindArmor(s)
	if i >= 0 {
		res := inv.Armors[i]
		inv.Armors = append(inv.Armors[0:i], inv.Armors[i+1:len(inv.Armors)]...)
		s.InventoryID = 0
		return res
	}
	return nil
}

// Omamoris

func (i *Inventory) AddOmamori(o *Omamori) {
	i.Omamoris = append(i.Omamoris, o)
}

func (inv *Inventory) FindOmamori(s *Omamori) int {
	for i, stbl := range inv.Omamoris {
		if stbl.ID == s.ID {
			return i
		}
	}
	return -1
}

func (inv *Inventory) RemoveOmamori(s *Omamori) *Omamori {
	i := inv.FindOmamori(s)
	if i >= 0 {
		res := inv.Omamoris[i]
		inv.Omamoris = append(inv.Omamoris[0:i], inv.Omamoris[i+1:len(inv.Omamoris)]...)
		s.InventoryID = 0
		return res
	}
	return nil
}

// Currencies

func (i *Inventory) AddCurrency(c *Currency, number int) {
	for _, co := range i.Currencies {
		if co.Currency.Type == c.Type {
			co.NumberOwned += number
			return
		}
	}
}

func (i *Inventory) FindCurrency(c *Currency) *CurrencyOwned {
	for _, co := range i.Currencies {
		if co.Currency.Type == c.Type {
			return co
		}
	}
	return nil
}

func (i *Inventory) RemoveCurrency(c *Currency, number int) int {
	co := i.FindCurrency(c)
	if co.NumberOwned >= number {
		co.NumberOwned -= number
		return number
	}
	return 0
}

// Inventory

func (i *Inventory) IsInInventory(s IStorable, number int) bool {
	switch st := s.(type) {
	case *Weapon:
		return (i.FindWeapon(st) >= 0)
	case *Armor:
		return (i.FindArmor(st) >= 0)
	case *Omamori:
		return (i.FindOmamori(st) >= 0)
	case *Currency:
		co := i.FindCurrency(st)
		return (co != nil && co.NumberOwned >= number)
	}
	return false
}

func (i *Inventory) HasEnought(s IStorable, number int) bool {
	switch st := s.(type) {
	case *Weapon:
		return len(i.Weapons) >= number
	case *Armor:
		return len(i.Armors) >= number
	case *Omamori:
		return len(i.Omamoris) >= number
	case *Currency:
		co := i.FindCurrency(st)
		return (co != nil && co.NumberOwned >= number)
	}
	return false
}

func (i *Inventory) Store(s IStorable, number int) {
	s.Store(i, number)
}

func (i *Inventory) Remove(s IStorable, number int) bool {
	if number == 0 {
		return true
	}
	switch st := s.(type) {
	case *Weapon:
		return i.RemoveWeapon(st) != nil
	case *Armor:
		return i.RemoveArmor(st) != nil
	case *Omamori:
		return i.RemoveOmamori(st) != nil
	case *Currency:
		return i.RemoveCurrency(st, number) == number
	}
	return false
}

func (inv *Inventory) Merge(other *Inventory) {
	inv.Weapons = append(inv.Weapons, other.Weapons...)
	inv.Armors = append(inv.Armors, other.Armors...)
	inv.Omamoris = append(inv.Omamoris, other.Omamoris...)
	for i := range inv.Currencies {
		inv.Currencies[i].NumberOwned += other.Currencies[i].NumberOwned
	}
}

func (i *Inventory) StoreReward(re []IStorable, cu map[CurrencyType]int) {
	for _, s := range re {
		s.Store(i, 1)
	}
	for k, v := range cu {
		i.Currencies[k].NumberOwned += v
	}
}
