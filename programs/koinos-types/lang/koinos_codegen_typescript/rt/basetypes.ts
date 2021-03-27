import * as ByteBuffer from 'bytebuffer'

export class KString {
  
  public _string: string;

  constructor(str: string = "") {
    this._string = str;
  }
  
  serialize(): VariableBlob {
    return new VariableBlob(ByteBuffer.fromUTF8(this._string));
  }
}

export class KBoolean {
  
  public _bool: boolean;

  constructor(bool: boolean = false) {
    this._bool = bool;
  }

  serialize(): VariableBlob {
    const buffer = new ByteBuffer();
    buffer.writeByte(this._bool ? 1: 0)
    return new VariableBlob(buffer);
  }
}

export class VariableBlob {

  public buffer: ByteBuffer;

  constructor(b: ByteBuffer | number = 0) {
    if(b instanceof ByteBuffer)
      this.buffer = b;
    else
      this.buffer = ByteBuffer.allocate(b);
  }

  serialize(): VariableBlob {
    if(this.buffer.offset !== 0) this.buffer.flip();
    const ser = new VariableBlob();
    ser.buffer
      .writeVarint64(this.buffer.limit)
      .append(this.buffer);
    return ser;
  }

  deserializeVariableBlob(): VariableBlob {
    if(this.buffer.offset !== 0) this.buffer.flip();
    const size = this.buffer.readVarint64().toNumber();
    if(size < 0)
      throw new Error("Could not deserialize variable blob");

    const { limit, offset } = this.buffer;
    if(limit < offset + size)
      throw new Error("Unexpected EOF");
    const vb = new VariableBlob(size);
    this.buffer.copyTo(vb.buffer, 0, offset, offset + size);
    return vb;
  }

  deserializeString(): KString {
    if(this.buffer.offset !== 0) this.buffer.flip();
    return new KString(this.buffer.toUTF8())
  }

  deserializeBoolean(): KBoolean {
    if(this.buffer.offset !== 0) this.buffer.flip();
    if(this.buffer.limit === 0)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readByte(0);
    if( value !== 0 && value !== 1)
      throw new Error("Boolean must be 0 or 1");
    return new KBoolean(!!value);
  }

  toHex() {
    if(this.buffer.offset !== 0) this.buffer.flip();
    return this.buffer.toHex();
  }
}
