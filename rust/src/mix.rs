#[allow(dead_code)]
pub fn str_example(){
    let mut str:String =  "hello".to_string();
    str.push_str(" ");
    str.push_str("World");
    print!("{}", str);
}


#[cfg(test)]
mod test{
    use super::str_example;

    #[test]
    fn    str_example_test(){
        str_example()
    }
}

