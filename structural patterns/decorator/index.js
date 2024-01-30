const ENV_ENCRYPTION_ALGORYTHM = "hgfads"
const ENV_COMPRESSION_ALGORYTHM = "20%"

class DataSource {
    readData() {
        throw new Error('Needs implementation!');
    }
    writeData() {
        throw new Error('Needs implementation!');
    }
}

class FileData extends DataSource {
    constructor(name, content) {
        super();
        this.name = name;
        if (content) this.writeData(content);
    }

    readData() {
        return this.content;
    }

    writeData(content) {
        this.content = content;
    }
}

class EncryptionDecorator extends DataSource {
    constructor(source) {
        super();
        this.source = source;
        this.encryptionAlgorithm = ENV_ENCRYPTION_ALGORYTHM
    }

    readData() {
        const data = this.source.readData();
        const regex = new RegExp(`${ENV_ENCRYPTION_ALGORYTHM}$`, "g");
        
        console.log("Decrypting data...");
        return data.replaceAll(regex, "");
    }

    writeData(data) {
        const encryData = (data || this.source.readData()) + ENV_ENCRYPTION_ALGORYTHM;
        
        console.log("Encrypting data...");
        this.source.writeData(encryData);
    }
}

class CompressionDecorator extends DataSource {
    constructor(source) {
        super();
        this.source = source;
        this.compressionAlgorithm = ENV_COMPRESSION_ALGORYTHM
    }

    readData() {
        const data = this.source.readData();
        const regex = new RegExp(`${ENV_COMPRESSION_ALGORYTHM}$`, "g");
        
        console.log("Decompressing data...");
        return data.replaceAll(regex, "");
    }

    writeData(data) {
        const comprData = (data || this.source.readData()) + ENV_COMPRESSION_ALGORYTHM;
        
        console.log("Compressing data...");
        this.source.writeData(comprData);
    }
}

let source = new FileData("someFile.txt", "Heya!");
console.log(source.readData());
source = new EncryptionDecorator(source);
source.writeData();
console.log(source.readData());
source = new CompressionDecorator(source);
source.writeData();
console.log(source.readData());