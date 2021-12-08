package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) int {
    if avg <= 0 {
        return len(layers) * 2
    } else {
        return len(layers) * avg
    }
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodl int, sauce float64) {
    for _, l := range layers {
        switch l {
        case "noodles":
            noodl += 50
        case "sauce":
            sauce += 0.2
        }
    }
    return noodl, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend, mine []string) (newOne []string) {
    for _, l := range mine {
        newOne = append(newOne, l)
    }
    newOne = append(newOne, friend[len(friend)-1])
    return newOne
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(qs []float64, portions int) (newOne []float64) {
    for _, q := range qs {
        newOne = append(newOne, (q/2)*float64(portions))
    }
    return newOne
}
