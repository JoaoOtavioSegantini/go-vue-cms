import Posts from "@/components/Posts.vue";
import { mount } from "@vue/test-utils";
import { createStore } from "vuex";

describe("Posts.vue", () => {
  const createVuexStore = () => {
    return createStore({
      state() {
        return {
          Posts: {
            all: [{ ID: "1", title: "Post 1", body: "A simple post body" }],
          },
        };
      },
      mutations: {
        updatePostsList(state) {
          state.Posts = [];
        },
      },
      actions: {
        listPosts(ctx) {
          const res = mockGet();
          ctx.commit("updatePostsList", res);
        },
      },
    });
  };

  function factory() {
    const store = createVuexStore();
    return mount(Posts, {
      global: {
        plugins: [store],
      },
    });
  }

  const mockGet = jest.fn();

  jest.mock("axios", () => ({
    get: () => mockGet(),
  }));

  it("renders props when passed", () => {
    const msg = "AÇÕES Artigos Gerenciamento de artigos do blog#titlebody";

    const wrapper = factory();

    expect(mockGet).toHaveBeenCalled();
    expect(wrapper.text()).toMatch(msg);
    expect(wrapper.find(".card").exists()).toBe(true)
    expect(wrapper.find(".card-body").exists()).toBe(true)
    expect(wrapper.find(".title").exists()).toBe(true)
    expect(wrapper.html()).toMatchSnapshot();
  });
});
